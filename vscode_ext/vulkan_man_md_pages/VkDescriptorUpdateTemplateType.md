# VkDescriptorUpdateTemplateType(3) Manual Page

## Name

VkDescriptorUpdateTemplateType - Indicates the valid usage of the
descriptor update template



## <a href="#_c_specification" class="anchor"></a>C Specification

The descriptor update template type is determined by the
[VkDescriptorUpdateTemplateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateCreateInfo.html)::`templateType`
property, which takes the following values:

``` c
// Provided by VK_VERSION_1_1
typedef enum VkDescriptorUpdateTemplateType {
    VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET = 0,
  // Provided by VK_VERSION_1_1 with VK_KHR_push_descriptor, VK_KHR_descriptor_update_template with VK_KHR_push_descriptor
    VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR = 1,
  // Provided by VK_KHR_descriptor_update_template
    VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET_KHR = VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET,
} VkDescriptorUpdateTemplateType;
```

or the equivalent

``` c
// Provided by VK_KHR_descriptor_update_template
typedef VkDescriptorUpdateTemplateType VkDescriptorUpdateTemplateTypeKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET` specifies that the
  descriptor update template will be used for descriptor set updates
  only.

- `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR` specifies
  that the descriptor update template will be used for push descriptor
  updates only.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkDescriptorUpdateTemplateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorUpdateTemplateType"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
