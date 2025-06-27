# VK_KHR_performance_query(3) Manual Page

## Name

VK_KHR_performance_query - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

117

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_special_use" class="anchor"></a>Special Use

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">Developer tools</a>

## <a href="#_contact" class="anchor"></a>Contact

- Alon Or-bach <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_performance_query%5D%20@alonorbach%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_performance_query%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>alonorbach</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

The `VK_KHR_performance_query` extension adds a mechanism to allow
querying of performance counters for use in applications and by
profiling tools.

Each queue family **may** expose counters that **can** be enabled on a
queue of that family. We extend [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html) to add a
new query type for performance queries, and chain a structure on
[VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html) to specify the
performance queries to enable.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkAcquireProfilingLockKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireProfilingLockKHR.html)

- [vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumeratePhysicalDeviceQueueFamilyPerformanceQueryCountersKHR.html)

- [vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyPerformanceQueryPassesKHR.html)

- [vkReleaseProfilingLockKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkReleaseProfilingLockKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkAcquireProfilingLockInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAcquireProfilingLockInfoKHR.html)

- [VkPerformanceCounterDescriptionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterDescriptionKHR.html)

- [VkPerformanceCounterKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterKHR.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDevicePerformanceQueryFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePerformanceQueryFeaturesKHR.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDevicePerformanceQueryPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePerformanceQueryPropertiesKHR.html)

- Extending [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html):

  - [VkQueryPoolPerformanceCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolPerformanceCreateInfoKHR.html)

- Extending [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html),
  [VkSubmitInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo2.html):

  - [VkPerformanceQuerySubmitInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceQuerySubmitInfoKHR.html)

## <a href="#_new_unions" class="anchor"></a>New Unions

- [VkPerformanceCounterResultKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterResultKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkAcquireProfilingLockFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAcquireProfilingLockFlagBitsKHR.html)

- [VkPerformanceCounterDescriptionFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterDescriptionFlagBitsKHR.html)

- [VkPerformanceCounterScopeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterScopeKHR.html)

- [VkPerformanceCounterStorageKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterStorageKHR.html)

- [VkPerformanceCounterUnitKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterUnitKHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkAcquireProfilingLockFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAcquireProfilingLockFlagsKHR.html)

- [VkPerformanceCounterDescriptionFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterDescriptionFlagsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_PERFORMANCE_QUERY_EXTENSION_NAME`

- `VK_KHR_PERFORMANCE_QUERY_SPEC_VERSION`

- Extending [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html):

  - `VK_QUERY_TYPE_PERFORMANCE_QUERY_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_ACQUIRE_PROFILING_LOCK_INFO_KHR`

  - `VK_STRUCTURE_TYPE_PERFORMANCE_COUNTER_DESCRIPTION_KHR`

  - `VK_STRUCTURE_TYPE_PERFORMANCE_COUNTER_KHR`

  - `VK_STRUCTURE_TYPE_PERFORMANCE_QUERY_SUBMIT_INFO_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PERFORMANCE_QUERY_FEATURES_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PERFORMANCE_QUERY_PROPERTIES_KHR`

  - `VK_STRUCTURE_TYPE_QUERY_POOL_PERFORMANCE_CREATE_INFO_KHR`

## <a href="#_issues" class="anchor"></a>Issues

1\) Should this extension include a mechanism to begin a query in
command buffer *A* and end the query in command buffer *B*?

**RESOLVED** No - queries are tied to command buffer creation and thus
have to be encapsulated within a single command buffer.

2\) Should this extension include a mechanism to begin and end queries
globally on the queue, not using the existing command buffer commands?

**RESOLVED** No - for the same reasoning as the resolution of 1).

3\) Should this extension expose counters that require multiple passes?

**RESOLVED** Yes - users should re-submit a command buffer with the same
commands in it multiple times, specifying the pass to count as the query
parameter in VkPerformanceQuerySubmitInfoKHR.

4\) How to handle counters across parallel workloads?

**RESOLVED** In the spirit of Vulkan, a counter description flag
`VK_PERFORMANCE_COUNTER_DESCRIPTION_CONCURRENTLY_IMPACTED_BIT_KHR`
denotes that the accuracy of a counter result is affected by parallel
workloads.

5\) How to handle secondary command buffers?

**RESOLVED** Secondary command buffers inherit any counter pass index
specified in the parent primary command buffer. Note: this is no longer
an issue after change from issue 10 resolution

6\) What commands does the profiling lock have to be held for?

**RESOLVED** For any command buffer that is being queried with a
performance query pool, the profiling lock **must** be held while that
command buffer is in the *recording*, *executable*, or *pending state*.

7\) Should we support
[vkCmdCopyQueryPoolResults](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyQueryPoolResults.html)?

**RESOLVED** Yes.

8\) Should we allow performance queries to interact with multiview?

**RESOLVED** Yes, but the performance queries must be performed once for
each pass per view.

9\) Should a `queryCount > 1` be usable for performance queries?

**RESOLVED** Yes. Some vendors will have costly performance counter
query pool creation, and would rather if a certain set of counters were
to be used multiple times that a `queryCount > 1` can be used to
amortize the instantiation cost.

10\) Should we introduce an indirect mechanism to set the counter pass
index?

**RESOLVED** Specify the counter pass index at submit time instead, to
avoid requiring re-recording of command buffers when multiple counter
passes are needed.

## <a href="#_examples" class="anchor"></a>Examples

The following example shows how to find what performance counters a
queue family supports, setup a query pool to record these performance
counters, how to add the query pool to the command buffer to record
information, and how to get the results from the query pool.

``` c
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

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-10-08

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_performance_query"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
