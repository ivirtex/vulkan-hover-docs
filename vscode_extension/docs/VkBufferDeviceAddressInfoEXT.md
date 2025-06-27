# VkBufferDeviceAddressInfo(3) Manual Page

## Name

VkBufferDeviceAddressInfo - Structure specifying the buffer to query an
address for



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBufferDeviceAddressInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkBufferDeviceAddressInfo {
    VkStructureType    sType;
    const void*        pNext;
    VkBuffer           buffer;
} VkBufferDeviceAddressInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_buffer_device_address
typedef VkBufferDeviceAddressInfo VkBufferDeviceAddressInfoKHR;
```

or the equivalent

``` c
// Provided by VK_EXT_buffer_device_address
typedef VkBufferDeviceAddressInfo VkBufferDeviceAddressInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `buffer` specifies the buffer whose address is being queried.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkBufferDeviceAddressInfo-buffer-02600"
  id="VUID-VkBufferDeviceAddressInfo-buffer-02600"></a>
  VUID-VkBufferDeviceAddressInfo-buffer-02600  
  If `buffer` is non-sparse and was not created with the
  `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` flag, then it
  **must** be bound completely and contiguously to a single
  `VkDeviceMemory` object

- <a href="#VUID-VkBufferDeviceAddressInfo-buffer-02601"
  id="VUID-VkBufferDeviceAddressInfo-buffer-02601"></a>
  VUID-VkBufferDeviceAddressInfo-buffer-02601  
  `buffer` **must** have been created with
  `VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT`

Valid Usage (Implicit)

- <a href="#VUID-VkBufferDeviceAddressInfo-sType-sType"
  id="VUID-VkBufferDeviceAddressInfo-sType-sType"></a>
  VUID-VkBufferDeviceAddressInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO`

- <a href="#VUID-VkBufferDeviceAddressInfo-pNext-pNext"
  id="VUID-VkBufferDeviceAddressInfo-pNext-pNext"></a>
  VUID-VkBufferDeviceAddressInfo-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkBufferDeviceAddressInfo-buffer-parameter"
  id="VUID-VkBufferDeviceAddressInfo-buffer-parameter"></a>
  VUID-VkBufferDeviceAddressInfo-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html),
[vkGetBufferDeviceAddressEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddressEXT.html),
[vkGetBufferDeviceAddressKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddressKHR.html),
[vkGetBufferOpaqueCaptureAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferOpaqueCaptureAddress.html),
[vkGetBufferOpaqueCaptureAddressKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferOpaqueCaptureAddressKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBufferDeviceAddressInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
