# VkPhysicalDevicePerStageDescriptorSetFeaturesNV(3) Manual Page

## Name

VkPhysicalDevicePerStageDescriptorSetFeaturesNV - Structure describing feature to allow descriptor set layout bindings to be per-stage



## [](#_c_specification)C Specification

The `VkPhysicalDevicePerStageDescriptorSetFeaturesNV` structure is defined as:

```c++
// Provided by VK_NV_per_stage_descriptor_set
typedef struct VkPhysicalDevicePerStageDescriptorSetFeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           perStageDescriptorSet;
    VkBool32           dynamicPipelineLayout;
} VkPhysicalDevicePerStageDescriptorSetFeaturesNV;
```

## [](#_members)Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`perStageDescriptorSet` indicates that the implementation allows the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PER_STAGE_BIT_NV` descriptor set layout creation flag to be used so the bindings are specified per-stage rather than across all stages.
- []()`dynamicPipelineLayout` indicates the implementation allows the `layout` member of [VkBindDescriptorSetsInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindDescriptorSetsInfo.html), [VkPushConstantsInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushConstantsInfo.html), [VkPushDescriptorSetInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushDescriptorSetInfo.html), [VkPushDescriptorSetWithTemplateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushDescriptorSetWithTemplateInfo.html), [VkSetDescriptorBufferOffsetsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSetDescriptorBufferOffsetsInfoEXT.html) and [VkBindDescriptorBufferEmbeddedSamplersInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindDescriptorBufferEmbeddedSamplersInfoEXT.html) to be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayoutCreateInfo.html) **can** be chained off those structures' `pNext` instead.

## [](#_description)Description

If the `VkPhysicalDevicePerStageDescriptorSetFeaturesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDevicePerStageDescriptorSetFeaturesNV`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDevicePerStageDescriptorSetFeaturesNV-sType-sType)VUID-VkPhysicalDevicePerStageDescriptorSetFeaturesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PER_STAGE_DESCRIPTOR_SET_FEATURES_NV`

## [](#_see_also)See Also

[VK\_NV\_per\_stage\_descriptor\_set](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_per_stage_descriptor_set.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDevicePerStageDescriptorSetFeaturesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0