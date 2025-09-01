# VkPhysicalDevicePipelineRobustnessProperties(3) Manual Page

## Name

VkPhysicalDevicePipelineRobustnessProperties - Structure describing the default robustness behavior of a physical device



## [](#_c_specification)C Specification

The `VkPhysicalDevicePipelineRobustnessProperties` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkPhysicalDevicePipelineRobustnessProperties {
    VkStructureType                       sType;
    void*                                 pNext;
    VkPipelineRobustnessBufferBehavior    defaultRobustnessStorageBuffers;
    VkPipelineRobustnessBufferBehavior    defaultRobustnessUniformBuffers;
    VkPipelineRobustnessBufferBehavior    defaultRobustnessVertexInputs;
    VkPipelineRobustnessImageBehavior     defaultRobustnessImages;
} VkPhysicalDevicePipelineRobustnessProperties;
```

or the equivalent

```c++
// Provided by VK_EXT_pipeline_robustness
typedef VkPhysicalDevicePipelineRobustnessProperties VkPhysicalDevicePipelineRobustnessPropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- `defaultRobustnessStorageBuffers` describes the behavior of out of bounds accesses made to storage buffers when no robustness features are enabled
- `defaultRobustnessUniformBuffers` describes the behavior of out of bounds accesses made to uniform buffers when no robustness features are enabled
- `defaultRobustnessVertexInputs` describes the behavior of out of bounds accesses made to vertex input attributes when no robustness features are enabled
- `defaultRobustnessImages` describes the behavior of out of bounds accesses made to images when no robustness features are enabled

Some implementations of Vulkan may be able to guarantee that certain types of accesses are always performed with robustness even when the Vulkan API’s robustness features are not explicitly enabled.

Even when an implementation reports that accesses to a given resource type are robust by default, it remains invalid to make an out of bounds access without requesting the appropriate robustness feature.

If the `VkPhysicalDevicePipelineRobustnessProperties` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDevicePipelineRobustnessProperties-sType-sType)VUID-VkPhysicalDevicePipelineRobustnessProperties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_ROBUSTNESS_PROPERTIES`

## [](#_see_also)See Also

[VK\_EXT\_pipeline\_robustness](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_pipeline_robustness.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkPipelineRobustnessBufferBehavior](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessBufferBehavior.html), [VkPipelineRobustnessImageBehavior](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessImageBehavior.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDevicePipelineRobustnessProperties).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0