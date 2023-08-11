import math 
import json 
import numpy as np
import matplotlib.pyplot as plt

# This function extracts the data from the json file and 
# returns it as an array. The values are in milliseconds.
def extract_data(filename):
    data = json.load(open(filename))
    times = []
    for value in data:
        time = value.get('Time')
        times.append(time/1000000)
    return times

# This is the plot for the sharing phase for all three sharing schemes. 
def runtime_sharing_all_schemes(): 
    # This is the values for n
    N = [3, 4, 5, 6, 7, 8, 9, 10, 11]
    X_axis = np.arange(len(N))

    # These are the measured times in ms for the secret-sharing schemes
    Shamir = np.array(extract_data("json/all/shamir_sharing.json"))   #[0.02318, 0.039086, 0.038589, 0.063819, 0.064952, 0.095104, 0.122583, 0.117813, 0.118148])
    Replicated = np.array(extract_data("json/all/replicated_sharing.json"))  #[1.998031, 3.271821, 6.213167, 13.595797, 27.677489, 109.107063, 787.750018, 6343.602573, 72901.618923])
    Additive = np.array(extract_data("json/all/additive_sharing.json"))  #[2.060047, 3.648013, 3.762626, 4.652979, 5.469777, 6.571979, 7.488585, 8.788275, 9.591049])

    # This is the plot
    plt.figure(figsize=(8, 4))

    plt.bar(X_axis - 0.2, Shamir, 0.2, color='#6b97aa', label='Shamir\'s secret-sharing')
    plt.bar(X_axis + 0., Replicated, 0.2, color='#154e56', label='Replicated secret-sharing')
    plt.bar(X_axis + 0.2, Additive, 0.2, color='#069668', label='Additive secret-sharing')

    plt.xticks(X_axis, N)

    plt.xlabel("Threshold")
    plt.ylabel("Running Time (ms)")

    plt.legend()

    plt.ylim([1,10**5])


    plt.yscale('log')


# This is the plot for the reconstruction phase for all three sharing schemes.
def runtime_reconstruction_all_schemes(): 
    # This is the values for n
    N = [3, 4, 5, 6, 7, 8, 9, 10, 11]
    X_axis = np.arange(len(N))

    # These are the measured times in ms for the secret-sharing schemes
    Shamir = np.array(extract_data("json/all/shamir_reconstruction.json"))
    Replicated = np.array(extract_data("json/all/replicated_reconstruction.json"))
    Additive = np.array(extract_data("json/all/additive_reconstruction.json"))

    # This is the plot
    plt.figure(figsize=(8, 4))

    plt.bar(X_axis - 0.2, Shamir, 0.2, color='#6b97aa', label='Shamir\'s secret-sharing')
    plt.bar(X_axis + 0., Replicated, 0.2, color='#154e56', label='Replicated secret-sharing')
    plt.bar(X_axis + 0.2, Additive, 0.2, color='#069668', label='Additive secret-sharing')

    plt.xticks(X_axis, N)

    plt.xlabel("Threshold")
    plt.ylabel("Running Time (ms)")

    plt.legend()

    plt.ylim([10**-3.5,10**0])


    plt.yscale('log')


# This is the plot for the number of shares for all three sharing schemes.
def shares_all_schemes(): 
    # This is the values for n
    N = [3, 4, 5, 6, 7, 8, 9, 10, 11]
    X_axis = np.arange(len(N))

    # These are the measured times in ms for the secret-sharing schemes
    Shamir = np.array([3, 4, 5, 6, 7, 8, 9, 10, 11])
    Replicated = np.array([3, 4, 10, 15, 35, 56, 126, 210, 462])
    Additive = np.array([3, 4, 5, 6, 7, 8, 9, 10, 11])

    # This is the plot
    plt.figure(figsize=(8, 4))

    plt.bar(X_axis - 0.2, Shamir, 0.2, color='#6b97aa', label='Shamir\'s secret-sharing')
    plt.bar(X_axis + 0., Replicated, 0.2, color='#154e56', label='Replicated secret-sharing')
    plt.bar(X_axis + 0.2, Additive, 0.2, color='#069668', label='Additive secret-sharing')

    plt.xticks(X_axis, N)

    plt.xlabel("Threshold")
    plt.ylabel("Number of shares")

    plt.legend()

    plt.ylim([1,10**3])

    plt.yscale('log')


# This is the plot for the memory usage for all three sharing schemes.
def memory_usage_all_schemes(): 
    # This is the values for n
    N = [3, 4, 5, 6, 7, 8, 9, 10, 11]
    X_axis = np.arange(len(N))

    # These are the measured times in ms for the secret-sharing schemes
    Shamir = np.array([4064, 4480, 8768, 9632, 17696, 19456, 30272, 32928, 46400])
    Replicated = np.array([334336, 502744, 1528206, 2639799, 8359644, 41648656, 352281024, 3203647056, 36762391376])
    Additive = np.array([322613, 484232, 656338, 813433, 971058, 1147520, 1296240, 1472363, 1634696])

    # This is the plot
    plt.figure(figsize=(8, 4))

    plt.bar(X_axis - 0.2, Shamir, 0.2, color='#6b97aa', label='Shamir\'s secret-sharing')
    plt.bar(X_axis + 0., Replicated, 0.2, color='#154e56', label='Replicated secret-sharing')
    plt.bar(X_axis + 0.2, Additive, 0.2, color='#069668', label='Additive secret-sharing')

    plt.xticks(X_axis, N)

    plt.xlabel("Threshold")
    plt.ylabel("Memory Usage (bytes)")

    plt.legend()

    plt.yscale('log')



if __name__ == "__main__":
    """
    runtime_sharing_all_schemes()
    plt.savefig("plots/all/runtime_sharing_all_schemes.pdf", bbox_inches='tight')
    plt.show()

    runtime_reconstruction_all_schemes()
    plt.savefig("plots/all/runtime_reconstruction_all_schemes.pdf", bbox_inches='tight')
    plt.show()
    """ 

    shares_all_schemes()
    plt.savefig("plots/all/shares_all_schemes.pdf", bbox_inches='tight')
    plt.show()

    memory_usage_all_schemes()
    plt.savefig("plots/all/memory_usage_all_schemes.pdf", bbox_inches='tight')
    plt.show()
