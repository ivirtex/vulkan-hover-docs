# vkGetDescriptorSetHostMappingVALVE(3) Manual Page

## Name

vkGetDescriptorSetHostMappingVALVE - Stub description of vkGetDescriptorSetHostMappingVALVE



## [](#_c_specification)C Specification

There is currently no specification language written for this command. This section acts only as placeholder and to avoid dead links in the specification and reference pages.

```c++
// Provided by VK_VALVE_descriptor_set_host_mapping
void vkGetDescriptorSetHostMappingVALVE(
    VkDevice                                    device,
    VkDescriptorSet                             descriptorSet,
    void**                                      ppData);
```

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetDescriptorSetHostMappingVALVE-device-parameter)VUID-vkGetDescriptorSetHostMappingVALVE-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDescriptorSetHostMappingVALVE-descriptorSet-parameter)VUID-vkGetDescriptorSetHostMappingVALVE-descriptorSet-parameter  
  `descriptorSet` **must** be a valid [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSet.html) handle
- [](#VUID-vkGetDescriptorSetHostMappingVALVE-ppData-parameter)VUID-vkGetDescriptorSetHostMappingVALVE-ppData-parameter  
  `ppData` **must** be a valid pointer to a pointer value
- [](#VUID-vkGetDescriptorSetHostMappingVALVE-descriptorSet-parent)VUID-vkGetDescriptorSetHostMappingVALVE-descriptorSet-parent  
  `descriptorSet` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_VALVE\_descriptor\_set\_host\_mapping](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VALVE_descriptor_set_host_mapping.html), [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSet.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDescriptorSetHostMappingVALVE).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0