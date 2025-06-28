# VK\_INTEL\_performance\_query(3) Manual Page

## Name

VK\_INTEL\_performance\_query - device extension



## [](#_registered_extension_number)Registered Extension Number

211

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_special_use)Special Use

- [Developer tools](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Lionel Landwerlin [\[GitHub\]llandwerlin](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_INTEL_performance_query%5D%20%40llandwerlin%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_INTEL_performance_query%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-05-16

**IP Status**

No known IP claims.

**Contributors**

- Lionel Landwerlin, Intel
- Piotr Maciejewski, Intel

## [](#_description)Description

This extension allows an application to capture performance data to be interpreted by an external application or library.

Such a library is available at : [https://github.com/intel/metrics-discovery](https://github.com/intel/metrics-discovery)

Performance analysis tools such as [Graphics Performance Analyzers](https://www.intel.com/content/www/us/en/developer/tools/graphics-performance-analyzers/overview.html) make use of this extension and the metrics-discovery library to present the data in a human readable way.

## [](#_new_object_types)New Object Types

- [VkPerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceConfigurationINTEL.html)

## [](#_new_commands)New Commands

- [vkAcquirePerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquirePerformanceConfigurationINTEL.html)
- [vkCmdSetPerformanceMarkerINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetPerformanceMarkerINTEL.html)
- [vkCmdSetPerformanceOverrideINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetPerformanceOverrideINTEL.html)
- [vkCmdSetPerformanceStreamMarkerINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetPerformanceStreamMarkerINTEL.html)
- [vkGetPerformanceParameterINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPerformanceParameterINTEL.html)
- [vkInitializePerformanceApiINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/vkInitializePerformanceApiINTEL.html)
- [vkQueueSetPerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueSetPerformanceConfigurationINTEL.html)
- [vkReleasePerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/vkReleasePerformanceConfigurationINTEL.html)
- [vkUninitializePerformanceApiINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUninitializePerformanceApiINTEL.html)

## [](#_new_structures)New Structures

- [VkInitializePerformanceApiInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInitializePerformanceApiInfoINTEL.html)
- [VkPerformanceConfigurationAcquireInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceConfigurationAcquireInfoINTEL.html)
- [VkPerformanceMarkerInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceMarkerInfoINTEL.html)
- [VkPerformanceOverrideInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceOverrideInfoINTEL.html)
- [VkPerformanceStreamMarkerInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceStreamMarkerInfoINTEL.html)
- [VkPerformanceValueINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceValueINTEL.html)
- Extending [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateInfo.html):
  
  - [VkQueryPoolCreateInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolCreateInfoINTEL.html)
  - [VkQueryPoolPerformanceQueryCreateInfoINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolPerformanceQueryCreateInfoINTEL.html)

## [](#_new_unions)New Unions

- [VkPerformanceValueDataINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceValueDataINTEL.html)

## [](#_new_enums)New Enums

- [VkPerformanceConfigurationTypeINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceConfigurationTypeINTEL.html)
- [VkPerformanceOverrideTypeINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceOverrideTypeINTEL.html)
- [VkPerformanceParameterTypeINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceParameterTypeINTEL.html)
- [VkPerformanceValueTypeINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceValueTypeINTEL.html)
- [VkQueryPoolSamplingModeINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPoolSamplingModeINTEL.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_INTEL_PERFORMANCE_QUERY_EXTENSION_NAME`
- `VK_INTEL_PERFORMANCE_QUERY_SPEC_VERSION`
- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html):
  
  - `VK_OBJECT_TYPE_PERFORMANCE_CONFIGURATION_INTEL`
- Extending [VkQueryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryType.html):
  
  - `VK_QUERY_TYPE_PERFORMANCE_QUERY_INTEL`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_INITIALIZE_PERFORMANCE_API_INFO_INTEL`
  - `VK_STRUCTURE_TYPE_PERFORMANCE_CONFIGURATION_ACQUIRE_INFO_INTEL`
  - `VK_STRUCTURE_TYPE_PERFORMANCE_MARKER_INFO_INTEL`
  - `VK_STRUCTURE_TYPE_PERFORMANCE_OVERRIDE_INFO_INTEL`
  - `VK_STRUCTURE_TYPE_PERFORMANCE_STREAM_MARKER_INFO_INTEL`
  - `VK_STRUCTURE_TYPE_QUERY_POOL_CREATE_INFO_INTEL`
  - `VK_STRUCTURE_TYPE_QUERY_POOL_PERFORMANCE_QUERY_CREATE_INFO_INTEL`

## [](#_example_code)Example Code

```c
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

## [](#_version_history)Version History

- Revision 2, 2020-03-06 (Lionel Landwerlin)
  
  - Rename VkQueryPoolCreateInfoINTEL in VkQueryPoolPerformanceQueryCreateInfoINTEL
- Revision 1, 2018-05-16 (Lionel Landwerlin)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_INTEL_performance_query)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0