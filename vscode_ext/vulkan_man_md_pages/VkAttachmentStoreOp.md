# VkAttachmentStoreOp(3) Manual Page

## Name

VkAttachmentStoreOp - Specify how contents of an attachment are treated
at the end of the subpass where it is last used



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription.html)::`storeOp` and
`stencilStoreOp`, specifying how the contents of the attachment are
treated, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkAttachmentStoreOp {
    VK_ATTACHMENT_STORE_OP_STORE = 0,
    VK_ATTACHMENT_STORE_OP_DONT_CARE = 1,
  // Provided by VK_VERSION_1_3
    VK_ATTACHMENT_STORE_OP_NONE = 1000301000,
  // Provided by VK_KHR_dynamic_rendering, VK_KHR_load_store_op_none
    VK_ATTACHMENT_STORE_OP_NONE_KHR = VK_ATTACHMENT_STORE_OP_NONE,
  // Provided by VK_QCOM_render_pass_store_ops
    VK_ATTACHMENT_STORE_OP_NONE_QCOM = VK_ATTACHMENT_STORE_OP_NONE,
  // Provided by VK_EXT_load_store_op_none
    VK_ATTACHMENT_STORE_OP_NONE_EXT = VK_ATTACHMENT_STORE_OP_NONE,
} VkAttachmentStoreOp;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_ATTACHMENT_STORE_OP_STORE` specifies the contents generated during
  the render pass and within the render area are written to memory. For
  attachments with a depth/stencil format, this uses the access type
  `VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT`. For attachments with a
  color format, this uses the access type
  `VK_ACCESS_COLOR_ATTACHMENT_WRITE_BIT`.

- `VK_ATTACHMENT_STORE_OP_DONT_CARE` specifies the contents within the
  render area are not needed after rendering, and **may** be discarded;
  the contents of the attachment will be undefined inside the render
  area. For attachments with a depth/stencil format, this uses the
  access type `VK_ACCESS_DEPTH_STENCIL_ATTACHMENT_WRITE_BIT`. For
  attachments with a color format, this uses the access type
  `VK_ACCESS_COLOR_ATTACHMENT_WRITE_BIT`.

- `VK_ATTACHMENT_STORE_OP_NONE` specifies the contents within the render
  area are not accessed by the store operation as long as no values are
  written to the attachment during the render pass. If values are
  written during the render pass, this behaves identically to
  `VK_ATTACHMENT_STORE_OP_DONT_CARE` and with matching access semantics.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p><code>VK_ATTACHMENT_STORE_OP_DONT_CARE</code> <strong>can</strong>
cause contents generated during previous render passes to be discarded
before reaching memory, even if no write to the attachment occurs during
the current render pass.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription.html),
[VkAttachmentDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription2.html),
[VkRenderingAttachmentInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAttachmentInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAttachmentStoreOp"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
