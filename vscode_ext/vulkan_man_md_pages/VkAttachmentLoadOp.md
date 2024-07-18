# VkAttachmentLoadOp(3) Manual Page

## Name

VkAttachmentLoadOp - Specify how contents of an attachment are treated
at the beginning of the subpass where it is first used



## <a href="#_c_specification" class="anchor"></a>C Specification

Load operations that **can** be used for a render pass are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkAttachmentLoadOp {
    VK_ATTACHMENT_LOAD_OP_LOAD = 0,
    VK_ATTACHMENT_LOAD_OP_CLEAR = 1,
    VK_ATTACHMENT_LOAD_OP_DONT_CARE = 2,
  // Provided by VK_KHR_load_store_op_none
    VK_ATTACHMENT_LOAD_OP_NONE_KHR = 1000400000,
  // Provided by VK_EXT_load_store_op_none
    VK_ATTACHMENT_LOAD_OP_NONE_EXT = VK_ATTACHMENT_LOAD_OP_NONE_KHR,
} VkAttachmentLoadOp;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_ATTACHMENT_LOAD_OP_LOAD` specifies that the previous contents of
  the image within the render area will be preserved as the initial
  values. For attachments with a depth/stencil format, this uses the
  access type `VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_READ_BIT`. For
  attachments with a color format, this uses the access type
  `VK_ACCESS_COLOR_ATTACHMENT_READ_BIT`.

- `VK_ATTACHMENT_LOAD_OP_CLEAR` specifies that the contents within the
  render area will be cleared to a uniform value, which is specified
  when a render pass instance is begun. For attachments with a
  depth/stencil format, this uses the access type
  `VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT`. For attachments with a
  color format, this uses the access type
  `VK_ACCESS_COLOR_ATTACHMENT_WRITE_BIT`.

- `VK_ATTACHMENT_LOAD_OP_DONT_CARE` specifies that the previous contents
  within the area need not be preserved; the contents of the attachment
  will be undefined inside the render area. For attachments with a
  depth/stencil format, this uses the access type
  `VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT`. For attachments with a
  color format, this uses the access type
  `VK_ACCESS_COLOR_ATTACHMENT_WRITE_BIT`.

- `VK_ATTACHMENT_LOAD_OP_NONE_KHR` specifies that the previous contents
  of the image will be undefined inside the render pass. No access type
  is used as the image is not accessed.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription.html),
[VkAttachmentDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription2.html),
[VkRenderingAttachmentInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAttachmentInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAttachmentLoadOp"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
