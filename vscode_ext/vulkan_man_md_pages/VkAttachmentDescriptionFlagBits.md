# VkAttachmentDescriptionFlagBits(3) Manual Page

## Name

VkAttachmentDescriptionFlagBits - Bitmask specifying additional
properties of an attachment



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription.html)::`flags`,
describing additional properties of the attachment, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkAttachmentDescriptionFlagBits {
    VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT = 0x00000001,
} VkAttachmentDescriptionFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT` specifies that the
  attachment aliases the same device memory as other attachments.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAttachmentDescriptionFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAttachmentDescriptionFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
