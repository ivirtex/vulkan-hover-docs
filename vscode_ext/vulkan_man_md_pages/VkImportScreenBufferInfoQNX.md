# VkImportScreenBufferInfoQNX(3) Manual Page

## Name

VkImportScreenBufferInfoQNX - Import memory from a QNX Screen buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To import memory created outside of the current Vulkan instance from a
QNX Screen buffer, add a `VkImportScreenBufferInfoQNX` structure to the
`pNext` chain of the [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html)
structure. The `VkImportScreenBufferInfoQNX` structure is defined as:

``` c
// Provided by VK_QNX_external_memory_screen_buffer
typedef struct VkImportScreenBufferInfoQNX {
    VkStructureType           sType;
    const void*               pNext;
    struct _screen_buffer*    buffer;
} VkImportScreenBufferInfoQNX;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `buffer` is a pointer to a `struct` `_screen_buffer`, the QNX Screen
  buffer to import

## <a href="#_description" class="anchor"></a>Description

The implementation **may** not acquire a reference to the imported
Screen buffer. Therefore, the application **must** ensure that the
object referred to by `buffer` stays valid as long as the device memory
to which it is imported is being used.

Valid Usage

- <a href="#VUID-VkImportScreenBufferInfoQNX-buffer-08966"
  id="VUID-VkImportScreenBufferInfoQNX-buffer-08966"></a>
  VUID-VkImportScreenBufferInfoQNX-buffer-08966  
  If `buffer` is not `NULL`, QNX Screen Buffers **must** be supported
  for import, as reported by
  [VkExternalImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalImageFormatProperties.html)
  or [VkExternalBufferProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalBufferProperties.html)

- <a href="#VUID-VkImportScreenBufferInfoQNX-buffer-08967"
  id="VUID-VkImportScreenBufferInfoQNX-buffer-08967"></a>
  VUID-VkImportScreenBufferInfoQNX-buffer-08967  
  `buffer` is not `NULL`, it **must** be a pointer to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-external-screen-buffer-validity"
  target="_blank" rel="noopener">valid QNX Screen buffer</a>

Valid Usage (Implicit)

- <a href="#VUID-VkImportScreenBufferInfoQNX-sType-sType"
  id="VUID-VkImportScreenBufferInfoQNX-sType-sType"></a>
  VUID-VkImportScreenBufferInfoQNX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMPORT_SCREEN_BUFFER_INFO_QNX`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QNX_external_memory_screen_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QNX_external_memory_screen_buffer.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImportScreenBufferInfoQNX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
