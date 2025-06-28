# VkExternalFormatANDROID(3) Manual Page

## Name

VkExternalFormatANDROID - Structure containing an Android hardware buffer external format



## [](#_c_specification)C Specification

`VkExternalFormatANDROID` is defined as:

```c++
// Provided by VK_ANDROID_external_memory_android_hardware_buffer
typedef struct VkExternalFormatANDROID {
    VkStructureType    sType;
    void*              pNext;
    uint64_t           externalFormat;
} VkExternalFormatANDROID;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `externalFormat` is an implementation-defined identifier for the external format

## [](#_description)Description

When included in the `pNext` chain of another structure, it indicates [additional format information](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-external-android-hardware-buffer-external-formats) beyond what is provided by `VkFormat` values for an Android hardware buffer. If `externalFormat` is zero, it indicates that no external format is used, and implementations should rely only on other format information. If this structure is not present, it is equivalent to setting `externalFormat` to zero.

Valid Usage

- [](#VUID-VkExternalFormatANDROID-externalFormat-01894)VUID-VkExternalFormatANDROID-externalFormat-01894  
  `externalFormat` **must** be `0` or a value returned in the `externalFormat` member of [VkAndroidHardwareBufferFormatPropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferFormatPropertiesANDROID.html) by an earlier call to [vkGetAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAndroidHardwareBufferPropertiesANDROID.html)

Valid Usage (Implicit)

- [](#VUID-VkExternalFormatANDROID-sType-sType)VUID-VkExternalFormatANDROID-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXTERNAL_FORMAT_ANDROID`

## [](#_see_also)See Also

[VK\_ANDROID\_external\_memory\_android\_hardware\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ANDROID_external_memory_android_hardware_buffer.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExternalFormatANDROID)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0