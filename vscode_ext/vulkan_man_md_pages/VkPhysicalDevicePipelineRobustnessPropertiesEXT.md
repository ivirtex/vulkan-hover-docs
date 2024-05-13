# VkPhysicalDevicePipelineRobustnessPropertiesEXT(3) Manual Page

## Name

VkPhysicalDevicePipelineRobustnessPropertiesEXT - Structure describing
the default robustness behavior of a physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDevicePipelineRobustnessPropertiesEXT` structure is
defined as:

``` c
// Provided by VK_EXT_pipeline_robustness
typedef struct VkPhysicalDevicePipelineRobustnessPropertiesEXT {
    VkStructureType                          sType;
    void*                                    pNext;
    VkPipelineRobustnessBufferBehaviorEXT    defaultRobustnessStorageBuffers;
    VkPipelineRobustnessBufferBehaviorEXT    defaultRobustnessUniformBuffers;
    VkPipelineRobustnessBufferBehaviorEXT    defaultRobustnessVertexInputs;
    VkPipelineRobustnessImageBehaviorEXT     defaultRobustnessImages;
} VkPhysicalDevicePipelineRobustnessPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `defaultRobustnessStorageBuffers` describes the behavior of out of
  bounds accesses made to storage buffers when no robustness features
  are enabled

- `defaultRobustnessUniformBuffers` describes the behavior of out of
  bounds accesses made to uniform buffers when no robustness features
  are enabled

- `defaultRobustnessVertexInputs` describes the behavior of out of
  bounds accesses made to vertex input attributes when no robustness
  features are enabled

- `defaultRobustnessImages` describes the behavior of out of bounds
  accesses made to images when no robustness features are enabled

## <a href="#_description" class="anchor"></a>Description

Some implementations of Vulkan may be able to guarantee that certain
types of accesses are always performed with robustness even when the
Vulkan API’s robustness features are not explicitly enabled.

Even when an implementation reports that accesses to a given resource
type are robust by default, it remains invalid to make an out of bounds
access without requesting the appropriate robustness feature.

If the `VkPhysicalDevicePipelineRobustnessPropertiesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDevicePipelineRobustnessPropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDevicePipelineRobustnessPropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDevicePipelineRobustnessPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_ROBUSTNESS_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_pipeline_robustness](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_pipeline_robustness.html),
[VkPipelineRobustnessBufferBehaviorEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRobustnessBufferBehaviorEXT.html),
[VkPipelineRobustnessImageBehaviorEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRobustnessImageBehaviorEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDevicePipelineRobustnessPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
