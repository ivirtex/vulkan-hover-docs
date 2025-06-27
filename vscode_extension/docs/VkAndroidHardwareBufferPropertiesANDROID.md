# VkAndroidHardwareBufferPropertiesANDROID(3) Manual Page

## Name

VkAndroidHardwareBufferPropertiesANDROID - Properties of External Memory
Android Hardware Buffers



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAndroidHardwareBufferPropertiesANDROID` structure returned is
defined as:

``` c
// Provided by VK_ANDROID_external_memory_android_hardware_buffer
typedef struct VkAndroidHardwareBufferPropertiesANDROID {
    VkStructureType    sType;
    void*              pNext;
    VkDeviceSize       allocationSize;
    uint32_t           memoryTypeBits;
} VkAndroidHardwareBufferPropertiesANDROID;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `allocationSize` is the size of the external memory

- `memoryTypeBits` is a bitmask containing one bit set for every memory
  type which the specified Android hardware buffer **can** be imported
  as.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkAndroidHardwareBufferPropertiesANDROID-sType-sType"
  id="VUID-VkAndroidHardwareBufferPropertiesANDROID-sType-sType"></a>
  VUID-VkAndroidHardwareBufferPropertiesANDROID-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_PROPERTIES_ANDROID`

- <a href="#VUID-VkAndroidHardwareBufferPropertiesANDROID-pNext-pNext"
  id="VUID-VkAndroidHardwareBufferPropertiesANDROID-pNext-pNext"></a>
  VUID-VkAndroidHardwareBufferPropertiesANDROID-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkAndroidHardwareBufferFormatProperties2ANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferFormatProperties2ANDROID.html),
  [VkAndroidHardwareBufferFormatPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferFormatPropertiesANDROID.html),
  or
  [VkAndroidHardwareBufferFormatResolvePropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferFormatResolvePropertiesANDROID.html)

- <a href="#VUID-VkAndroidHardwareBufferPropertiesANDROID-sType-unique"
  id="VUID-VkAndroidHardwareBufferPropertiesANDROID-sType-unique"></a>
  VUID-VkAndroidHardwareBufferPropertiesANDROID-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

## <a href="#_see_also" class="anchor"></a>See Also

[VK_ANDROID_external_memory_android_hardware_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ANDROID_external_memory_android_hardware_buffer.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAndroidHardwareBufferPropertiesANDROID.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAndroidHardwareBufferPropertiesANDROID"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
