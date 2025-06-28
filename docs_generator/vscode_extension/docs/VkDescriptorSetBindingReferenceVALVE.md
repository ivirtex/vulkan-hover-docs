# VkDescriptorSetBindingReferenceVALVE(3) Manual Page

## Name

VkDescriptorSetBindingReferenceVALVE - Stub description of VkDescriptorSetBindingReferenceVALVE



## [](#_c_specification)C Specification

There is currently no specification language written for this type. This section acts only as placeholder and to avoid dead links in the specification and reference pages.

```c++
// Provided by VK_VALVE_descriptor_set_host_mapping
typedef struct VkDescriptorSetBindingReferenceVALVE {
    VkStructureType          sType;
    const void*              pNext;
    VkDescriptorSetLayout    descriptorSetLayout;
    uint32_t                 binding;
} VkDescriptorSetBindingReferenceVALVE;
```

## [](#_members)Members

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDescriptorSetBindingReferenceVALVE-sType-sType)VUID-VkDescriptorSetBindingReferenceVALVE-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_BINDING_REFERENCE_VALVE`
- [](#VUID-VkDescriptorSetBindingReferenceVALVE-pNext-pNext)VUID-VkDescriptorSetBindingReferenceVALVE-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDescriptorSetBindingReferenceVALVE-descriptorSetLayout-parameter)VUID-VkDescriptorSetBindingReferenceVALVE-descriptorSetLayout-parameter  
  `descriptorSetLayout` **must** be a valid [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html) handle

## [](#_see_also)See Also

[VK\_VALVE\_descriptor\_set\_host\_mapping](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VALVE_descriptor_set_host_mapping.html), [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetDescriptorSetLayoutHostMappingInfoVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDescriptorSetLayoutHostMappingInfoVALVE.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorSetBindingReferenceVALVE)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0