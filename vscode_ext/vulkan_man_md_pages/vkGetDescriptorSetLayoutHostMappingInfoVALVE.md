# vkGetDescriptorSetLayoutHostMappingInfoVALVE(3) Manual Page

## Name

vkGetDescriptorSetLayoutHostMappingInfoVALVE - Stub description of
vkGetDescriptorSetLayoutHostMappingInfoVALVE



## <a href="#_c_specification" class="anchor"></a>C Specification

There is currently no specification language written for this command.
This section acts only as placeholder and to avoid dead links in the
specification and reference pages.

``` c
// Provided by VK_VALVE_descriptor_set_host_mapping
void vkGetDescriptorSetLayoutHostMappingInfoVALVE(
    VkDevice                                    device,
    const VkDescriptorSetBindingReferenceVALVE* pBindingReference,
    VkDescriptorSetLayoutHostMappingInfoVALVE*  pHostMapping);
```

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetDescriptorSetLayoutHostMappingInfoVALVE-device-parameter"
  id="VUID-vkGetDescriptorSetLayoutHostMappingInfoVALVE-device-parameter"></a>
  VUID-vkGetDescriptorSetLayoutHostMappingInfoVALVE-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetDescriptorSetLayoutHostMappingInfoVALVE-pBindingReference-parameter"
  id="VUID-vkGetDescriptorSetLayoutHostMappingInfoVALVE-pBindingReference-parameter"></a>
  VUID-vkGetDescriptorSetLayoutHostMappingInfoVALVE-pBindingReference-parameter  
  `pBindingReference` **must** be a valid pointer to a valid
  [VkDescriptorSetBindingReferenceVALVE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetBindingReferenceVALVE.html)
  structure

- <a
  href="#VUID-vkGetDescriptorSetLayoutHostMappingInfoVALVE-pHostMapping-parameter"
  id="VUID-vkGetDescriptorSetLayoutHostMappingInfoVALVE-pHostMapping-parameter"></a>
  VUID-vkGetDescriptorSetLayoutHostMappingInfoVALVE-pHostMapping-parameter  
  `pHostMapping` **must** be a valid pointer to a
  [VkDescriptorSetLayoutHostMappingInfoVALVE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutHostMappingInfoVALVE.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VALVE_descriptor_set_host_mapping](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VALVE_descriptor_set_host_mapping.html),
[VkDescriptorSetBindingReferenceVALVE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetBindingReferenceVALVE.html),
[VkDescriptorSetLayoutHostMappingInfoVALVE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutHostMappingInfoVALVE.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDescriptorSetLayoutHostMappingInfoVALVE"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
