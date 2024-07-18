# vkGetDescriptorSetHostMappingVALVE(3) Manual Page

## Name

vkGetDescriptorSetHostMappingVALVE - Stub description of
vkGetDescriptorSetHostMappingVALVE



## <a href="#_c_specification" class="anchor"></a>C Specification

There is currently no specification language written for this command.
This section acts only as placeholder and to avoid dead links in the
specification and reference pages.

``` c
// Provided by VK_VALVE_descriptor_set_host_mapping
void vkGetDescriptorSetHostMappingVALVE(
    VkDevice                                    device,
    VkDescriptorSet                             descriptorSet,
    void**                                      ppData);
```

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkGetDescriptorSetHostMappingVALVE-device-parameter"
  id="VUID-vkGetDescriptorSetHostMappingVALVE-device-parameter"></a>
  VUID-vkGetDescriptorSetHostMappingVALVE-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetDescriptorSetHostMappingVALVE-descriptorSet-parameter"
  id="VUID-vkGetDescriptorSetHostMappingVALVE-descriptorSet-parameter"></a>
  VUID-vkGetDescriptorSetHostMappingVALVE-descriptorSet-parameter  
  `descriptorSet` **must** be a valid
  [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSet.html) handle

- <a href="#VUID-vkGetDescriptorSetHostMappingVALVE-ppData-parameter"
  id="VUID-vkGetDescriptorSetHostMappingVALVE-ppData-parameter"></a>
  VUID-vkGetDescriptorSetHostMappingVALVE-ppData-parameter  
  `ppData` **must** be a valid pointer to a pointer value

- <a href="#VUID-vkGetDescriptorSetHostMappingVALVE-descriptorSet-parent"
  id="VUID-vkGetDescriptorSetHostMappingVALVE-descriptorSet-parent"></a>
  VUID-vkGetDescriptorSetHostMappingVALVE-descriptorSet-parent  
  `descriptorSet` **must** have been created, allocated, or retrieved
  from `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VALVE_descriptor_set_host_mapping](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VALVE_descriptor_set_host_mapping.html),
[VkDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSet.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDescriptorSetHostMappingVALVE"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
