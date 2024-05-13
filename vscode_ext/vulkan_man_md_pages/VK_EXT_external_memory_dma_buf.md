# VK_EXT_external_memory_dma_buf(3) Manual Page

## Name

VK_EXT_external_memory_dma_buf - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

126

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_external_memory_fd](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory_fd.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Lina Versace <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_external_memory_dma_buf%5D%20@versalinyaa%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_external_memory_dma_buf%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>versalinyaa</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-10-10

**IP Status**  
No known IP claims.

**Contributors**  
- Lina Versace, Google

- James Jones, NVIDIA

- Faith Ekstrand, Intel

## <a href="#_description" class="anchor"></a>Description

A `dma_buf` is a type of file descriptor, defined by the Linux kernel,
that allows sharing memory across kernel device drivers and across
processes. This extension enables applications to import a `dma_buf` as
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html), to export
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) as a `dma_buf`, and to create
[VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) objects that **can** be bound to that memory.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_EXTERNAL_MEMORY_DMA_BUF_EXTENSION_NAME`

- `VK_EXT_EXTERNAL_MEMORY_DMA_BUF_SPEC_VERSION`

- Extending
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html):

  - `VK_EXTERNAL_MEMORY_HANDLE_TYPE_DMA_BUF_BIT_EXT`

## <a href="#_issues" class="anchor"></a>Issues

1\) How does the application, when creating a [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html)
that it intends to bind to `dma_buf`
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) containing an externally produced
image, specify the memory layout (such as row pitch and DRM format
modifier) of the [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html)? In other words, how does the
application achieve behavior comparable to that provided by
[`EGL_EXT_image_dma_buf_import`](https://registry.khronos.org/EGL/extensions/EXT/EGL_EXT_image_dma_buf_import.txt)
and
[`EGL_EXT_image_dma_buf_import_modifiers`](https://registry.khronos.org/EGL/extensions/EXT/EGL_EXT_image_dma_buf_import_modifiers.txt)
?

**RESOLVED**: Features comparable to those in
[`EGL_EXT_image_dma_buf_import`](https://registry.khronos.org/EGL/extensions/EXT/EGL_EXT_image_dma_buf_import.txt)
and
[`EGL_EXT_image_dma_buf_import_modifiers`](https://registry.khronos.org/EGL/extensions/EXT/EGL_EXT_image_dma_buf_import_modifiers.txt)
will be provided by an extension layered atop this one.

2\) Without the ability to specify the memory layout of external
`dma_buf` images, how is this extension useful?

**RESOLVED**: This extension provides exactly one new feature: the
ability to import/export between `dma_buf` and
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html). This feature, together with
features provided by
[`VK_KHR_external_memory_fd`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory_fd.html), is
sufficient to bind a [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) to `dma_buf`.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-10-10 (Lina Versace)

  - Squashed internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_external_memory_dma_buf"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
