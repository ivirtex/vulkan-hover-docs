# VkExternalFormatANDROID(3) Manual Page

## Name

VkExternalFormatANDROID - Structure containing an Android hardware
buffer external format



## <a href="#_c_specification" class="anchor"></a>C Specification

`VkExternalFormatANDROID` is defined as:

``` c
// Provided by VK_ANDROID_external_memory_android_hardware_buffer
typedef struct VkExternalFormatANDROID {
    VkStructureType    sType;
    void*              pNext;
    uint64_t           externalFormat;
} VkExternalFormatANDROID;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `externalFormat` is an implementation-defined identifier for the
  external format

## <a href="#_description" class="anchor"></a>Description

When included in the `pNext` chain of another structure, it indicates <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-external-android-hardware-buffer-external-formats"
target="_blank" rel="noopener">additional format information</a> beyond
what is provided by `VkFormat` values for an Android hardware buffer. If
`externalFormat` is zero, it indicates that no external format is used,
and implementations should rely only on other format information. If
this structure is not present, it is equivalent to setting
`externalFormat` to zero.

Valid Usage

- <a href="#VUID-VkExternalFormatANDROID-externalFormat-01894"
  id="VUID-VkExternalFormatANDROID-externalFormat-01894"></a>
  VUID-VkExternalFormatANDROID-externalFormat-01894  
  `externalFormat` **must** be `0` or a value returned in the
  `externalFormat` member of
  [VkAndroidHardwareBufferFormatPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferFormatPropertiesANDROID.html)
  by an earlier call to
  [vkGetAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAndroidHardwareBufferPropertiesANDROID.html)

Valid Usage (Implicit)

- <a href="#VUID-VkExternalFormatANDROID-sType-sType"
  id="VUID-VkExternalFormatANDROID-sType-sType"></a>
  VUID-VkExternalFormatANDROID-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXTERNAL_FORMAT_ANDROID`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_ANDROID_external_memory_android_hardware_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ANDROID_external_memory_android_hardware_buffer.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExternalFormatANDROID"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
