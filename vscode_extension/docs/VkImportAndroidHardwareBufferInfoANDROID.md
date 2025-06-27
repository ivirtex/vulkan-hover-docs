# VkImportAndroidHardwareBufferInfoANDROID(3) Manual Page

## Name

VkImportAndroidHardwareBufferInfoANDROID - Import memory from an Android
hardware buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To import memory created outside of the current Vulkan instance from an
Android hardware buffer, add a
`VkImportAndroidHardwareBufferInfoANDROID` structure to the `pNext`
chain of the [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html)
structure. The `VkImportAndroidHardwareBufferInfoANDROID` structure is
defined as:

``` c
// Provided by VK_ANDROID_external_memory_android_hardware_buffer
typedef struct VkImportAndroidHardwareBufferInfoANDROID {
    VkStructureType            sType;
    const void*                pNext;
    struct AHardwareBuffer*    buffer;
} VkImportAndroidHardwareBufferInfoANDROID;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `buffer` is the Android hardware buffer to import.

## <a href="#_description" class="anchor"></a>Description

If the [vkAllocateMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAllocateMemory.html) command succeeds, the
implementation **must** acquire a reference to the imported hardware
buffer, which it **must** release when the device memory object is
freed. If the command fails, the implementation **must** not retain a
reference.

Valid Usage

- <a href="#VUID-VkImportAndroidHardwareBufferInfoANDROID-buffer-01880"
  id="VUID-VkImportAndroidHardwareBufferInfoANDROID-buffer-01880"></a>
  VUID-VkImportAndroidHardwareBufferInfoANDROID-buffer-01880  
  If `buffer` is not `NULL`, Android hardware buffers **must** be
  supported for import, as reported by
  [VkExternalImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalImageFormatProperties.html)
  or [VkExternalBufferProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalBufferProperties.html)

- <a href="#VUID-VkImportAndroidHardwareBufferInfoANDROID-buffer-01881"
  id="VUID-VkImportAndroidHardwareBufferInfoANDROID-buffer-01881"></a>
  VUID-VkImportAndroidHardwareBufferInfoANDROID-buffer-01881  
  If `buffer` is not `NULL`, it **must** be a valid Android hardware
  buffer object with `AHardwareBuffer_Desc`::`usage` compatible with
  Vulkan as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-external-android-hardware-buffer"
  target="_blank" rel="noopener">Android Hardware Buffers</a>

Valid Usage (Implicit)

- <a href="#VUID-VkImportAndroidHardwareBufferInfoANDROID-sType-sType"
  id="VUID-VkImportAndroidHardwareBufferInfoANDROID-sType-sType"></a>
  VUID-VkImportAndroidHardwareBufferInfoANDROID-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMPORT_ANDROID_HARDWARE_BUFFER_INFO_ANDROID`

- <a
  href="#VUID-VkImportAndroidHardwareBufferInfoANDROID-buffer-parameter"
  id="VUID-VkImportAndroidHardwareBufferInfoANDROID-buffer-parameter"></a>
  VUID-VkImportAndroidHardwareBufferInfoANDROID-buffer-parameter  
  `buffer` **must** be a valid pointer to an
  [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/AHardwareBuffer.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_ANDROID_external_memory_android_hardware_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ANDROID_external_memory_android_hardware_buffer.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImportAndroidHardwareBufferInfoANDROID"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
