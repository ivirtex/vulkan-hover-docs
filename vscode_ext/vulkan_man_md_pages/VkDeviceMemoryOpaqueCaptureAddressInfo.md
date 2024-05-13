# VkDeviceMemoryOpaqueCaptureAddressInfo(3) Manual Page

## Name

VkDeviceMemoryOpaqueCaptureAddressInfo - Structure specifying the memory
object to query an address for



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDeviceMemoryOpaqueCaptureAddressInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkDeviceMemoryOpaqueCaptureAddressInfo {
    VkStructureType    sType;
    const void*        pNext;
    VkDeviceMemory     memory;
} VkDeviceMemoryOpaqueCaptureAddressInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_buffer_device_address
typedef VkDeviceMemoryOpaqueCaptureAddressInfo VkDeviceMemoryOpaqueCaptureAddressInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `memory` specifies the memory whose address is being queried.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkDeviceMemoryOpaqueCaptureAddressInfo-memory-03336"
  id="VUID-VkDeviceMemoryOpaqueCaptureAddressInfo-memory-03336"></a>
  VUID-VkDeviceMemoryOpaqueCaptureAddressInfo-memory-03336  
  `memory` **must** have been allocated with
  `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT`

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceMemoryOpaqueCaptureAddressInfo-sType-sType"
  id="VUID-VkDeviceMemoryOpaqueCaptureAddressInfo-sType-sType"></a>
  VUID-VkDeviceMemoryOpaqueCaptureAddressInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEVICE_MEMORY_OPAQUE_CAPTURE_ADDRESS_INFO`

- <a href="#VUID-VkDeviceMemoryOpaqueCaptureAddressInfo-pNext-pNext"
  id="VUID-VkDeviceMemoryOpaqueCaptureAddressInfo-pNext-pNext"></a>
  VUID-VkDeviceMemoryOpaqueCaptureAddressInfo-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkDeviceMemoryOpaqueCaptureAddressInfo-memory-parameter"
  id="VUID-VkDeviceMemoryOpaqueCaptureAddressInfo-memory-parameter"></a>
  VUID-VkDeviceMemoryOpaqueCaptureAddressInfo-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_buffer_device_address](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_buffer_device_address.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetDeviceMemoryOpaqueCaptureAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceMemoryOpaqueCaptureAddress.html),
[vkGetDeviceMemoryOpaqueCaptureAddressKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceMemoryOpaqueCaptureAddressKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceMemoryOpaqueCaptureAddressInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
