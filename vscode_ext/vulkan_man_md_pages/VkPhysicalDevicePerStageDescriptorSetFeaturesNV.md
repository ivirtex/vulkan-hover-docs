# VkPhysicalDevicePerStageDescriptorSetFeaturesNV(3) Manual Page

## Name

VkPhysicalDevicePerStageDescriptorSetFeaturesNV - Structure describing
feature to allow descriptor set layout bindings to be per-stage



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDevicePerStageDescriptorSetFeaturesNV` structure is
defined as:

``` c
// Provided by VK_NV_per_stage_descriptor_set
typedef struct VkPhysicalDevicePerStageDescriptorSetFeaturesNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           perStageDescriptorSet;
    VkBool32           dynamicPipelineLayout;
} VkPhysicalDevicePerStageDescriptorSetFeaturesNV;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-perStageDescriptorSet"></span>
  `perStageDescriptorSet` indicates that the implementation allows the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PER_STAGE_BIT_NV` descriptor set
  layout creation flag to be used so the bindings are specified
  per-stage rather than across all stages.

- <span id="features-dynamicPipelineLayout"></span>
  `dynamicPipelineLayout` indicates the implementation allows the
  `layout` member of
  [VkBindDescriptorSetsInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindDescriptorSetsInfoKHR.html),
  [VkPushConstantsInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushConstantsInfoKHR.html),
  [VkPushDescriptorSetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushDescriptorSetInfoKHR.html),
  [VkPushDescriptorSetWithTemplateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushDescriptorSetWithTemplateInfoKHR.html),
  [VkSetDescriptorBufferOffsetsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSetDescriptorBufferOffsetsInfoEXT.html)
  and
  [VkBindDescriptorBufferEmbeddedSamplersInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindDescriptorBufferEmbeddedSamplersInfoEXT.html)
  to be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html) **can**
  be chained off those structures' `pNext` instead.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDevicePerStageDescriptorSetFeaturesNV` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDevicePerStageDescriptorSetFeaturesNV` **can** also be used
in the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDevicePerStageDescriptorSetFeaturesNV-sType-sType"
  id="VUID-VkPhysicalDevicePerStageDescriptorSetFeaturesNV-sType-sType"></a>
  VUID-VkPhysicalDevicePerStageDescriptorSetFeaturesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PER_STAGE_DESCRIPTOR_SET_FEATURES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_per_stage_descriptor_set](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_per_stage_descriptor_set.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDevicePerStageDescriptorSetFeaturesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
