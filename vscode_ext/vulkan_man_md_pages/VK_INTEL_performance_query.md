# VK_INTEL_performance_query(3) Manual Page

## Name

VK_INTEL_performance_query - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

211

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_special_use" class="anchor"></a>Special Use

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">Developer tools</a>

## <a href="#_contact" class="anchor"></a>Contact

- Lionel Landwerlin <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_INTEL_performance_query%5D%20@llandwerlin%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_INTEL_performance_query%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>llandwerlin</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-05-16

**IP Status**  
No known IP claims.

**Contributors**  
- Lionel Landwerlin, Intel

- Piotr Maciejewski, Intel

## <a href="#_description" class="anchor"></a>Description

This extension allows an application to capture performance data to be
interpreted by an external application or library.

Such a library is available at :
<a href="https://github.com/intel/metrics-discovery"
class="bare">https://github.com/intel/metrics-discovery</a>

Performance analysis tools such as [Graphics Performance
Analyzers](https://software.intel.com/content/www/us/en/develop/tools/graphics-performance-analyzers.html)
make use of this extension and the metrics-discovery library to present
the data in a human readable way.

## <a href="#_new_object_types" class="anchor"></a>New Object Types

- [VkPerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceConfigurationINTEL.html)

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkAcquirePerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquirePerformanceConfigurationINTEL.html)

- [vkCmdSetPerformanceMarkerINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPerformanceMarkerINTEL.html)

- [vkCmdSetPerformanceOverrideINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPerformanceOverrideINTEL.html)

- [vkCmdSetPerformanceStreamMarkerINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPerformanceStreamMarkerINTEL.html)

- [vkGetPerformanceParameterINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPerformanceParameterINTEL.html)

- [vkInitializePerformanceApiINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkInitializePerformanceApiINTEL.html)

- [vkQueueSetPerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueSetPerformanceConfigurationINTEL.html)

- [vkReleasePerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkReleasePerformanceConfigurationINTEL.html)

- [vkUninitializePerformanceApiINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUninitializePerformanceApiINTEL.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkInitializePerformanceApiInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInitializePerformanceApiInfoINTEL.html)

- [VkPerformanceConfigurationAcquireInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceConfigurationAcquireInfoINTEL.html)

- [VkPerformanceMarkerInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceMarkerInfoINTEL.html)

- [VkPerformanceOverrideInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceOverrideInfoINTEL.html)

- [VkPerformanceStreamMarkerInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceStreamMarkerInfoINTEL.html)

- [VkPerformanceValueINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceValueINTEL.html)

- Extending [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html):

  - [VkQueryPoolCreateInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfoINTEL.html)

  - [VkQueryPoolPerformanceQueryCreateInfoINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolPerformanceQueryCreateInfoINTEL.html)

## <a href="#_new_unions" class="anchor"></a>New Unions

- [VkPerformanceValueDataINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceValueDataINTEL.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkPerformanceConfigurationTypeINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceConfigurationTypeINTEL.html)

- [VkPerformanceOverrideTypeINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceOverrideTypeINTEL.html)

- [VkPerformanceParameterTypeINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceParameterTypeINTEL.html)

- [VkPerformanceValueTypeINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceValueTypeINTEL.html)

- [VkQueryPoolSamplingModeINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolSamplingModeINTEL.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_INTEL_PERFORMANCE_QUERY_EXTENSION_NAME`

- `VK_INTEL_PERFORMANCE_QUERY_SPEC_VERSION`

- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html):

  - `VK_OBJECT_TYPE_PERFORMANCE_CONFIGURATION_INTEL`

- Extending [VkQueryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryType.html):

  - `VK_QUERY_TYPE_PERFORMANCE_QUERY_INTEL`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_INITIALIZE_PERFORMANCE_API_INFO_INTEL`

  - `VK_STRUCTURE_TYPE_PERFORMANCE_CONFIGURATION_ACQUIRE_INFO_INTEL`

  - `VK_STRUCTURE_TYPE_PERFORMANCE_MARKER_INFO_INTEL`

  - `VK_STRUCTURE_TYPE_PERFORMANCE_OVERRIDE_INFO_INTEL`

  - `VK_STRUCTURE_TYPE_PERFORMANCE_STREAM_MARKER_INFO_INTEL`

  - `VK_STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO_INTEL`

  - `VK_STRUCTURE_TYPE_QUERY_POOL_PERFORMANCE_QUERY_CREATE_INFO_INTEL`

## <a href="#_example_code" class="anchor"></a>Example Code

``` c
// A previously created device
VkDevice device;

// A queue derived from the device
VkQueue queue;

VkInitializePerformanceApiInfoINTEL performanceApiInfoIntel = {
  VK_STRUCTURE_TYPE_INITIALIZE_PERFORMANCE_API_INFO_INTEL,
  NULL,
  NULL
};

vkInitializePerformanceApiINTEL(
  device,
  &performanceApiInfoIntel);

VkQueryPoolPerformanceQueryCreateInfoINTEL queryPoolIntel = {
  VK_STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO_INTEL,
  NULL,
  VK_QUERY_POOL_SAMPLING_MODE_MANUAL_INTEL,
};

VkQueryPoolCreateInfo queryPoolCreateInfo = {
  VK_STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO,
  &queryPoolIntel,
  0,
  VK_QUERY_TYPE_PERFORMANCE_QUERY_INTEL,
  1,
  0
};

VkQueryPool queryPool;

VkResult result = vkCreateQueryPool(
  device,
  &queryPoolCreateInfo,
  NULL,
  &queryPool);

assert(VK_SUCCESS == result);

// A command buffer we want to record counters on
VkCommandBuffer commandBuffer;

VkCommandBufferBeginInfo commandBufferBeginInfo = {
  VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO,
  NULL,
  VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT,
  NULL
};

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

VkPerformanceConfigurationAcquireInfoINTEL performanceConfigurationAcquireInfo = {
  VK_STRUCTURE_TYPE_PERFORMANCE_CONFIGURATION_ACQUIRE_INFO_INTEL,
  NULL,
  VK_PERFORMANCE_CONFIGURATION_TYPE_COMMAND_QUEUE_METRICS_DISCOVERY_ACTIVATED_INTEL
};

VkPerformanceConfigurationINTEL performanceConfigurationIntel;

result = vkAcquirePerformanceConfigurationINTEL(
  device,
  &performanceConfigurationAcquireInfo,
  &performanceConfigurationIntel);

vkQueueSetPerformanceConfigurationINTEL(queue, performanceConfigurationIntel);

assert(VK_SUCCESS == result);

// Submit the command buffer and wait for its completion
// ...

result = vkReleasePerformanceConfigurationINTEL(
  device,
  performanceConfigurationIntel);

assert(VK_SUCCESS == result);

// Get the report size from metrics-discovery's QueryReportSize

result = vkGetQueryPoolResults(
  device,
  queryPool,
  0, 1, QueryReportSize,
  data, QueryReportSize, 0);

assert(VK_SUCCESS == result);

// The data can then be passed back to metrics-discovery from which
// human readable values can be queried.
```

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2020-03-06 (Lionel Landwerlin)

  - Rename VkQueryPoolCreateInfoINTEL in
    VkQueryPoolPerformanceQueryCreateInfoINTEL

- Revision 1, 2018-05-16 (Lionel Landwerlin)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_INTEL_performance_query"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
