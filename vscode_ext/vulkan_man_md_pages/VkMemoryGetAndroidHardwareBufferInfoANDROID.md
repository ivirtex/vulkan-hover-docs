# VkMemoryGetAndroidHardwareBufferInfoANDROID(3) Manual Page

## Name

VkMemoryGetAndroidHardwareBufferInfoANDROID - Structure describing an
Android hardware buffer memory export operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMemoryGetAndroidHardwareBufferInfoANDROID` structure is defined
as:

``` c
// Provided by VK_ANDROID_external_memory_android_hardware_buffer
typedef struct VkMemoryGetAndroidHardwareBufferInfoANDROID {
    VkStructureType    sType;
    const void*        pNext;
    VkDeviceMemory     memory;
} VkMemoryGetAndroidHardwareBufferInfoANDROID;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `memory` is the memory object from which the Android hardware buffer
  will be exported.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-handleTypes-01882"
  id="VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-handleTypes-01882"></a>
  VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-handleTypes-01882  
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID`
  **must** have been included in
  [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfo.html)::`handleTypes`
  when `memory` was created

- <a href="#VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-pNext-01883"
  id="VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-pNext-01883"></a>
  VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-pNext-01883  
  If the `pNext` chain of the
  [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) used to allocate
  `memory` included a
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)
  with non-`NULL` `image` member, then that `image` **must** already be
  bound to `memory`

Valid Usage (Implicit)

- <a href="#VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-sType-sType"
  id="VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-sType-sType"></a>
  VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_MEMORY_GET_ANDROID_HARDWARE_BUFFER_INFO_ANDROID`

- <a href="#VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-pNext-pNext"
  id="VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-pNext-pNext"></a>
  VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-pNext-pNext  
  `pNext` **must** be `NULL`

- <a
  href="#VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-memory-parameter"
  id="VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-memory-parameter"></a>
  VUID-VkMemoryGetAndroidHardwareBufferInfoANDROID-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_ANDROID_external_memory_android_hardware_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ANDROID_external_memory_android_hardware_buffer.html),
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetMemoryAndroidHardwareBufferANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryAndroidHardwareBufferANDROID.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryGetAndroidHardwareBufferInfoANDROID"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
