# VkDescriptorSetBindingReferenceVALVE(3) Manual Page

## Name

VkDescriptorSetBindingReferenceVALVE - Stub description of
VkDescriptorSetBindingReferenceVALVE



## <a href="#_c_specification" class="anchor"></a>C Specification

There is currently no specification language written for this type. This
section acts only as placeholder and to avoid dead links in the
specification and reference pages.

``` c
// Provided by VK_VALVE_descriptor_set_host_mapping
typedef struct VkDescriptorSetBindingReferenceVALVE {
    VkStructureType          sType;
    const void*              pNext;
    VkDescriptorSetLayout    descriptorSetLayout;
    uint32_t                 binding;
} VkDescriptorSetBindingReferenceVALVE;
```

## <a href="#_members" class="anchor"></a>Members

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkDescriptorSetBindingReferenceVALVE-sType-sType"
  id="VUID-VkDescriptorSetBindingReferenceVALVE-sType-sType"></a>
  VUID-VkDescriptorSetBindingReferenceVALVE-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_BINDING_REFERENCE_VALVE`

- <a href="#VUID-VkDescriptorSetBindingReferenceVALVE-pNext-pNext"
  id="VUID-VkDescriptorSetBindingReferenceVALVE-pNext-pNext"></a>
  VUID-VkDescriptorSetBindingReferenceVALVE-pNext-pNext  
  `pNext` **must** be `NULL`

- <a
  href="#VUID-VkDescriptorSetBindingReferenceVALVE-descriptorSetLayout-parameter"
  id="VUID-VkDescriptorSetBindingReferenceVALVE-descriptorSetLayout-parameter"></a>
  VUID-VkDescriptorSetBindingReferenceVALVE-descriptorSetLayout-parameter  
  `descriptorSetLayout` **must** be a valid
  [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VALVE_descriptor_set_host_mapping](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VALVE_descriptor_set_host_mapping.html),
[VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetDescriptorSetLayoutHostMappingInfoVALVE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorSetLayoutHostMappingInfoVALVE.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorSetBindingReferenceVALVE"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
