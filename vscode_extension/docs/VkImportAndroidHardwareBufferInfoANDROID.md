# VkImportAndroidHardwareBufferInfoANDROID(3) Manual Page

## Name

VkImportAndroidHardwareBufferInfoANDROID - Import memory from an Android hardware buffer



## [](#_c_specification)C Specification

To import memory created outside of the current Vulkan instance from an Android hardware buffer, add a `VkImportAndroidHardwareBufferInfoANDROID` structure to the `pNext` chain of the [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) structure. The `VkImportAndroidHardwareBufferInfoANDROID` structure is defined as:

```c++
// Provided by VK_ANDROID_external_memory_android_hardware_buffer
typedef struct VkImportAndroidHardwareBufferInfoANDROID {
    VkStructureType            sType;
    const void*                pNext;
    struct AHardwareBuffer*    buffer;
} VkImportAndroidHardwareBufferInfoANDROID;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `buffer` is the Android hardware buffer to import.

## [](#_description)Description

If the [vkAllocateMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAllocateMemory.html) command succeeds, the implementation **must** acquire a reference to the imported hardware buffer, which it **must** release when the device memory object is freed. If the command fails, the implementation **must** not retain a reference.

Valid Usage

- [](#VUID-VkImportAndroidHardwareBufferInfoANDROID-buffer-09863)VUID-VkImportAndroidHardwareBufferInfoANDROID-buffer-09863  
  If `buffer` is not `NULL`, Android hardware buffers **must** be supported for import, as reported by [VkExternalTensorPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalTensorPropertiesARM.html), [VkExternalImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalImageFormatProperties.html), or [VkExternalBufferProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalBufferProperties.html)
- [](#VUID-VkImportAndroidHardwareBufferInfoANDROID-buffer-01881)VUID-VkImportAndroidHardwareBufferInfoANDROID-buffer-01881  
  If `buffer` is not `NULL`, it **must** be a valid Android hardware buffer object with `AHardwareBuffer_Desc`::`usage` compatible with Vulkan as described in [Android Hardware Buffers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-external-android-hardware-buffer)

Valid Usage (Implicit)

- [](#VUID-VkImportAndroidHardwareBufferInfoANDROID-sType-sType)VUID-VkImportAndroidHardwareBufferInfoANDROID-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMPORT_ANDROID_HARDWARE_BUFFER_INFO_ANDROID`
- [](#VUID-VkImportAndroidHardwareBufferInfoANDROID-buffer-parameter)VUID-VkImportAndroidHardwareBufferInfoANDROID-buffer-parameter  
  `buffer` **must** be a valid pointer to an [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/AHardwareBuffer.html) value

## [](#_see_also)See Also

[VK\_ANDROID\_external\_memory\_android\_hardware\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ANDROID_external_memory_android_hardware_buffer.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImportAndroidHardwareBufferInfoANDROID)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0