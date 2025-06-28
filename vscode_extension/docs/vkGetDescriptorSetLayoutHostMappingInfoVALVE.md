# vkGetDescriptorSetLayoutHostMappingInfoVALVE(3) Manual Page

## Name

vkGetDescriptorSetLayoutHostMappingInfoVALVE - Stub description of vkGetDescriptorSetLayoutHostMappingInfoVALVE



## [](#_c_specification)C Specification

There is currently no specification language written for this command. This section acts only as placeholder and to avoid dead links in the specification and reference pages.

```c++
// Provided by VK_VALVE_descriptor_set_host_mapping
void vkGetDescriptorSetLayoutHostMappingInfoVALVE(
    VkDevice                                    device,
    const VkDescriptorSetBindingReferenceVALVE* pBindingReference,
    VkDescriptorSetLayoutHostMappingInfoVALVE*  pHostMapping);
```

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetDescriptorSetLayoutHostMappingInfoVALVE-device-parameter)VUID-vkGetDescriptorSetLayoutHostMappingInfoVALVE-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDescriptorSetLayoutHostMappingInfoVALVE-pBindingReference-parameter)VUID-vkGetDescriptorSetLayoutHostMappingInfoVALVE-pBindingReference-parameter  
  `pBindingReference` **must** be a valid pointer to a valid [VkDescriptorSetBindingReferenceVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetBindingReferenceVALVE.html) structure
- [](#VUID-vkGetDescriptorSetLayoutHostMappingInfoVALVE-pHostMapping-parameter)VUID-vkGetDescriptorSetLayoutHostMappingInfoVALVE-pHostMapping-parameter  
  `pHostMapping` **must** be a valid pointer to a [VkDescriptorSetLayoutHostMappingInfoVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutHostMappingInfoVALVE.html) structure

## [](#_see_also)See Also

[VK\_VALVE\_descriptor\_set\_host\_mapping](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VALVE_descriptor_set_host_mapping.html), [VkDescriptorSetBindingReferenceVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetBindingReferenceVALVE.html), [VkDescriptorSetLayoutHostMappingInfoVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutHostMappingInfoVALVE.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDescriptorSetLayoutHostMappingInfoVALVE)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0