# VK\_EXT\_external\_memory\_dma\_buf(3) Manual Page

## Name

VK\_EXT\_external\_memory\_dma\_buf - device extension



## [](#_registered_extension_number)Registered Extension Number

126

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_external\_memory\_fd](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory_fd.html)

## [](#_contact)Contact

- Lina Versace [\[GitHub\]linyaa-kiwi](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_external_memory_dma_buf%5D%20%40linyaa-kiwi%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_external_memory_dma_buf%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-10-10

**IP Status**

No known IP claims.

**Contributors**

- Lina Versace, Google
- James Jones, NVIDIA
- Faith Ekstrand, Intel

## [](#_description)Description

A `dma_buf` is a type of file descriptor, defined by the Linux kernel, that allows sharing memory across kernel device drivers and across processes. This extension enables applications to import a `dma_buf` as [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), to export [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) as a `dma_buf`, and to create [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) objects that **can** be bound to that memory.

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_EXTERNAL_MEMORY_DMA_BUF_EXTENSION_NAME`
- `VK_EXT_EXTERNAL_MEMORY_DMA_BUF_SPEC_VERSION`
- Extending [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html):
  
  - `VK_EXTERNAL_MEMORY_HANDLE_TYPE_DMA_BUF_BIT_EXT`

## [](#_issues)Issues

1\) How does the application, when creating a [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) that it intends to bind to `dma_buf` [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) containing an externally produced image, specify the memory layout (such as row pitch and DRM format modifier) of the [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html)? In other words, how does the application achieve behavior comparable to that provided by [`EGL_EXT_image_dma_buf_import`](https://registry.khronos.org/EGL/extensions/EXT/EGL_EXT_image_dma_buf_import.txt) and [`EGL_EXT_image_dma_buf_import_modifiers`](https://registry.khronos.org/EGL/extensions/EXT/EGL_EXT_image_dma_buf_import_modifiers.txt) ?

**RESOLVED**: Features comparable to those in [`EGL_EXT_image_dma_buf_import`](https://registry.khronos.org/EGL/extensions/EXT/EGL_EXT_image_dma_buf_import.txt) and [`EGL_EXT_image_dma_buf_import_modifiers`](https://registry.khronos.org/EGL/extensions/EXT/EGL_EXT_image_dma_buf_import_modifiers.txt) will be provided by an extension layered atop this one.

2\) Without the ability to specify the memory layout of external `dma_buf` images, how is this extension useful?

**RESOLVED**: This extension provides exactly one new feature: the ability to import/export between `dma_buf` and [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html). This feature, together with features provided by `VK_KHR_external_memory_fd`, is sufficient to bind a [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) to `dma_buf`.

## [](#_version_history)Version History

- Revision 1, 2017-10-10 (Lina Versace)
  
  - Squashed internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_external_memory_dma_buf).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0