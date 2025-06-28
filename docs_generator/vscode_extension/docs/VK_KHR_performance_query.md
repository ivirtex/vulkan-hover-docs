# VK\_KHR\_performance\_query(3) Manual Page

## Name

VK\_KHR\_performance\_query - device extension



## [](#_registered_extension_number)Registered Extension Number

117

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_special_use)Special Use

- [Developer tools](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Alon Or-bach [\[GitHub\]alonorbach](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_performance_query%5D%20%40alonorbach%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_performance_query%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-10-08

**IP Status**

No known IP claims.

**Contributors**

- Jesse Barker, Unity Technologies
- Kenneth Benzie, Codeplay
- Jan-Harald Fredriksen, ARM
- Jeff Leger, Qualcomm
- Jesse Hall, Google
- Tobias Hector, AMD
- Neil Henning, Codeplay
- Baldur Karlsson
- Lionel Landwerlin, Intel
- Peter Lohrmann, AMD
- Alon Or-bach, Samsung
- Daniel Rakos, AMD
- Niklas Smedberg, Unity Technologies
- Igor Ostrowski, Intel

## [](#_description)Description

The `VK_KHR_performance_query` extension adds a mechanism to allow querying of performance counters for use in applications and by profiling tools.

Each queue family **may** expose counters that **can** be enabled on a queue of that family. We extend [VkQueryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryType.html) to add a new query type for performance queries, and chain a structure on [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateInfo.html) to specify the performance queries to enable.

## [](#_new_commands)New Commands

- [vkAcquireProfilingLockKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireProfilingLockKHR.html)
- [vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR.html)
- [vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR.html)
- [vkReleaseProfilingLockKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkReleaseProfilingLockKHR.html)

## [](#_new_structures)New Structures

- [VkAcquireProfilingLockInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAcquireProfilingLockInfoKHR.html)
- [VkPerformanceCounterDescriptionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterDescriptionKHR.html)
- [VkPerformanceCounterKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterKHR.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDevicePerformanceQueryFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePerformanceQueryFeaturesKHR.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDevicePerformanceQueryPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePerformanceQueryPropertiesKHR.html)
- Extending [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateInfo.html):
  
  - [VkQueryPoolPerformanceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolPerformanceCreateInfoKHR.html)
- Extending [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html), [VkSubmitInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo2.html):
  
  - [VkPerformanceQuerySubmitInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceQuerySubmitInfoKHR.html)

## [](#_new_unions)New Unions

- [VkPerformanceCounterResultKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterResultKHR.html)

## [](#_new_enums)New Enums

- [VkAcquireProfilingLockFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAcquireProfilingLockFlagBitsKHR.html)
- [VkPerformanceCounterDescriptionFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterDescriptionFlagBitsKHR.html)
- [VkPerformanceCounterScopeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterScopeKHR.html)
- [VkPerformanceCounterStorageKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterStorageKHR.html)
- [VkPerformanceCounterUnitKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterUnitKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkAcquireProfilingLockFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAcquireProfilingLockFlagsKHR.html)
- [VkPerformanceCounterDescriptionFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceCounterDescriptionFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_PERFORMANCE_QUERY_EXTENSION_NAME`
- `VK_KHR_PERFORMANCE_QUERY_SPEC_VERSION`
- Extending [VkQueryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryType.html):
  
  - `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_ACQUIRE_PROFILING_LOCK_INFO_KHR`
  - `VK_STRUCTURE_TYPE_PERFORMANCE_COUNTER_DESCRIPTION_KHR`
  - `VK_STRUCTURE_TYPE_PERFORMANCE_COUNTER_KHR`
  - `VK_STRUCTURE_TYPE_PERFORMANCE_QUERY_SUBMIT_INFO_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PERFORMANCE_QUERY_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PERFORMANCE_QUERY_PROPERTIES_KHR`
  - `VK_STRUCTURE_TYPE_QUERY_POOL_PERFORMANCE_CREATE_INFO_KHR`

## [](#_issues)Issues

1\) Should this extension include a mechanism to begin a query in command buffer *A* and end the query in command buffer *B*?

**RESOLVED** No - queries are tied to command buffer creation and thus have to be encapsulated within a single command buffer.

2\) Should this extension include a mechanism to begin and end queries globally on the queue, not using the existing command buffer commands?

**RESOLVED** No - for the same reasoning as the resolution of 1).

3\) Should this extension expose counters that require multiple passes?

**RESOLVED** Yes - users should re-submit a command buffer with the same commands in it multiple times, specifying the pass to count as the query parameter in VkPerformanceQuerySubmitInfoKHR.

4\) How to handle counters across parallel workloads?

**RESOLVED** In the spirit of Vulkan, a counter description flag `VK_PERFORMANCE_COUNTER_DESCRIPTION_CONCURRENTLY_IMPACTED_BIT_KHR` denotes that the accuracy of a counter result is affected by parallel workloads.

5\) How to handle secondary command buffers?

**RESOLVED** Secondary command buffers inherit any counter pass index specified in the parent primary command buffer. Note: this is no longer an issue after change from issue 10 resolution

6\) What commands does the profiling lock have to be held for?

**RESOLVED** For any command buffer that is being queried with a performance query pool, the profiling lock **must** be held while that command buffer is in the *recording*, *executable*, or *pending state*.

7\) Should we support [vkCmdCopyQueryPoolResults](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyQueryPoolResults.html)?

**RESOLVED** Yes.

8\) Should we allow performance queries to interact with multiview?

**RESOLVED** Yes, but the performance queries must be performed once for each pass per view.

9\) Should a `queryCount > 1` be usable for performance queries?

**RESOLVED** Yes. Some vendors will have costly performance counter query pool creation, and would rather if a certain set of counters were to be used multiple times that a `queryCount > 1` can be used to amortize the instantiation cost.

10\) Should we introduce an indirect mechanism to set the counter pass index?

**RESOLVED** Specify the counter pass index at submit time instead, to avoid requiring re-recording of command buffers when multiple counter passes are needed.

## [](#_examples)Examples

The following example shows how to find what performance counters a queue family supports, setup a query pool to record these performance counters, how to add the query pool to the command buffer to record information, and how to get the results from the query pool.

```c++
// A previously created physical device
VkPhysicalDevice physicalDevice;

// One of the queue families our device supports
uint32_t queueFamilyIndex;

uint32_t counterCount;

// Get the count of counters supported
vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR(
  physicalDevice,
  queueFamilyIndex,
  &counterCount,
  NULL,
  NULL);

VkPerformanceCounterKHR* counters =
  malloc(sizeof(VkPerformanceCounterKHR) * counterCount);
VkPerformanceCounterDescriptionKHR* counterDescriptions =
  malloc(sizeof(VkPerformanceCounterDescriptionKHR) * counterCount);

// Get the counters supported
vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR(
  physicalDevice,
  queueFamilyIndex,
  &counterCount,
  counters,
  counterDescriptions);

// Try to enable the first 8 counters
uint32_t enabledCounters[8];

const uint32_t enabledCounterCount = min(counterCount, 8));

for (uint32_t i = 0; i < enabledCounterCount; i++) {
  enabledCounters[i] = i;
}

// A previously created device that had the performanceCounterQueryPools feature
// set to VK_TRUE
VkDevice device;

VkQueryPoolPerformanceCreateInfoKHR performanceQueryCreateInfo = {
  .sType = VK_STRUCTURE_TYPE_QUERY_POOL_PERFORMANCE_CREATE_INFO_KHR,
  .pNext = NULL,

  // Specify the queue family that this performance query is performed on
  .queueFamilyIndex = queueFamilyIndex,

  // The number of counters to enable
  .counterIndexCount = enabledCounterCount,

  // The array of indices of counters to enable
  .pCounterIndices = enabledCounters
};


// Get the number of passes our counters will require.
uint32_t numPasses;

vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR(
  physicalDevice,
  &performanceQueryCreateInfo,
  &numPasses);

VkQueryPoolCreateInfo queryPoolCreateInfo = {
  .sType = VK_STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO,
  .pNext = &performanceQueryCreateInfo,
  .flags = 0,
  // Using our new query type here
  .queryType = VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR,
  .queryCount = 1,
  .pipelineStatistics = 0
};

VkQueryPool queryPool;

VkResult result = vkCreateQueryPool(
  device,
  &queryPoolCreateInfo,
  NULL,
  &queryPool);

assert(VK_SUCCESS == result);

// A queue from queueFamilyIndex
VkQueue queue;

// A command buffer we want to record counters on
VkCommandBuffer commandBuffer;

VkCommandBufferBeginInfo commandBufferBeginInfo = {
  .sType = VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO,
  .pNext = NULL,
  .flags = 0,
  .pInheritanceInfo = NULL
};

VkAcquireProfilingLockInfoKHR lockInfo = {
  .sType = VK_STRUCTURE_TYPE_ACQUIRE_PROFILING_LOCK_INFO_KHR,
  .pNext = NULL,
  .flags = 0,
  .timeout = UINT64_MAX // Wait forever for the lock
};

// Acquire the profiling lock before we record command buffers
// that will use performance queries

result = vkAcquireProfilingLockKHR(device, &lockInfo);

assert(VK_SUCCESS == result);

result = vkBeginCommandBuffer(commandBuffer, &commandBufferBeginInfo);

assert(VK_SUCCESS == result);

vkCmdResetQueryPool(
  commandBuffer,
  queryPool,
  0,
  1);

vkCmdBeginQuery(
  commandBuffer,
  queryPool,
  0,
  0);

// Perform the commands you want to get performance information on
// ...

// Perform a barrier to ensure all previous commands were complete before
// ending the query
vkCmdPipelineBarrier(commandBuffer,
  VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT,
  VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT,
  0,
  0,
  NULL,
  0,
  NULL,
  0,
  NULL);

vkCmdEndQuery(
  commandBuffer,
  queryPool,
  0);

result = vkEndCommandBuffer(commandBuffer);

assert(VK_SUCCESS == result);

for (uint32_t counterPass = 0; counterPass < numPasses; counterPass++) {

  VkPerformanceQuerySubmitInfoKHR performanceQuerySubmitInfo = {
    VK_STRUCTURE_TYPE_PERFORMANCE_QUERY_SUBMIT_INFO_KHR,
    NULL,
    counterPass
  };


  // Submit the command buffer and wait for its completion
  // ...
}

// Release the profiling lock after the command buffer is no longer in the
// pending state.
vkReleaseProfilingLockKHR(device);

result = vkResetCommandBuffer(commandBuffer, 0);

assert(VK_SUCCESS == result);

// Create an array to hold the results of all counters
VkPerformanceCounterResultKHR* recordedCounters = malloc(
  sizeof(VkPerformanceCounterResultKHR) * enabledCounterCount);

result = vkGetQueryPoolResults(
  device,
  queryPool,
  0,
  1,
  sizeof(VkPerformanceCounterResultKHR) * enabledCounterCount,
  recordedCounters,
  sizeof(VkPerformanceCounterResultKHR) * enabledCounterCount,
  NULL);

// recordedCounters is filled with our counters, we will look at one for posterity
switch (counters[0].storage) {
  case VK_PERFORMANCE_COUNTER_STORAGE_INT32:
    // use recordCounters[0].int32 to get at the counter result!
    break;
  case VK_PERFORMANCE_COUNTER_STORAGE_INT64:
    // use recordCounters[0].int64 to get at the counter result!
    break;
  case VK_PERFORMANCE_COUNTER_STORAGE_UINT32:
    // use recordCounters[0].uint32 to get at the counter result!
    break;
  case VK_PERFORMANCE_COUNTER_STORAGE_UINT64:
    // use recordCounters[0].uint64 to get at the counter result!
    break;
  case VK_PERFORMANCE_COUNTER_STORAGE_FLOAT32:
    // use recordCounters[0].float32 to get at the counter result!
    break;
  case VK_PERFORMANCE_COUNTER_STORAGE_FLOAT64:
    // use recordCounters[0].float64 to get at the counter result!
    break;
}
```

## [](#_version_history)Version History

- Revision 1, 2019-10-08

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_performance_query)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0