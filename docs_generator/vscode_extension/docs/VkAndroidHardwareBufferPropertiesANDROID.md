# VkAndroidHardwareBufferPropertiesANDROID(3) Manual Page

## Name

VkAndroidHardwareBufferPropertiesANDROID - Properties of External Memory Android Hardware Buffers



## [](#_c_specification)C Specification

The `VkAndroidHardwareBufferPropertiesANDROID` structure returned is defined as:

```c++
// Provided by VK_ANDROID_external_memory_android_hardware_buffer
typedef struct VkAndroidHardwareBufferPropertiesANDROID {
    VkStructureType    sType;
    void*              pNext;
    VkDeviceSize       allocationSize;
    uint32_t           memoryTypeBits;
} VkAndroidHardwareBufferPropertiesANDROID;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `allocationSize` is the size of the external memory
- `memoryTypeBits` is a bitmask containing one bit set for every memory type which the specified Android hardware buffer **can** be imported as.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkAndroidHardwareBufferPropertiesANDROID-sType-sType)VUID-VkAndroidHardwareBufferPropertiesANDROID-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_PROPERTIES_ANDROID`
- [](#VUID-VkAndroidHardwareBufferPropertiesANDROID-pNext-pNext)VUID-VkAndroidHardwareBufferPropertiesANDROID-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkAndroidHardwareBufferFormatProperties2ANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferFormatProperties2ANDROID.html), [VkAndroidHardwareBufferFormatPropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferFormatPropertiesANDROID.html), or [VkAndroidHardwareBufferFormatResolvePropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferFormatResolvePropertiesANDROID.html)
- [](#VUID-VkAndroidHardwareBufferPropertiesANDROID-sType-unique)VUID-VkAndroidHardwareBufferPropertiesANDROID-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique

## [](#_see_also)See Also

[VK\_ANDROID\_external\_memory\_android\_hardware\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ANDROID_external_memory_android_hardware_buffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAndroidHardwareBufferPropertiesANDROID.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAndroidHardwareBufferPropertiesANDROID)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0