# VkMemoryGetAndroidHardwareBufferInfoANDROID(3) Manual Page

## Name

VkMemoryGetAndroidHardwareBufferInfoANDROID - Structure describing an Android hardware buffer memory export operation



## [](#_c_specification)C Specification

The `VkMemoryGetAndroidHardwareBufferInfoANDROID` structure is defined as:

```c++
// Provided by VK_ANDROID_external_memory_android_hardware_buffer
typedef struct VkMemoryGetAndroidHardwareBufferInfoANDROID {
    VkStructureType    sType;
    const void*        pNext;
    VkDeviceMemory     memory;
} VkMemoryGetAndroidHardwareBufferInfoANDROID;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `memory` is the memory object from which the Android hardware buffer will be exported.

## [](#_description)Description

Valid Usage

- [](#VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-handleTypes-01882)VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-handleTypes-01882  
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID` **must** have been included in [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMemoryAllocateInfo.html)::`handleTypes` when `memory` was created
- [](#VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-pNext-01883)VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-pNext-01883  
  If the `pNext` chain of the [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) used to allocate `memory` included a [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDedicatedAllocateInfo.html) with non-`NULL` `image` member, then that `image` **must** already be bound to `memory`

Valid Usage (Implicit)

- [](#VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-sType-sType)VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_GET_ANDROID_HARDWARE_BUFFER_INFO_ANDROID`
- [](#VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-pNext-pNext)VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-memory-parameter)VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle

## [](#_see_also)See Also

[VK\_ANDROID\_external\_memory\_android\_hardware\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ANDROID_external_memory_android_hardware_buffer.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetMemoryAndroidHardwareBufferANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryAndroidHardwareBufferANDROID.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryGetAndroidHardwareBufferInfoANDROID).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0