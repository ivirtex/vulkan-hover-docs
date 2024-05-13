# VkMemoryAllocateFlagBits(3) Manual Page

## Name

VkMemoryAllocateFlagBits - Bitmask specifying flags for a device memory
allocation



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkMemoryAllocateFlagsInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagsInfo.html)::`flags`,
controlling device memory allocation, are:

``` c
// Provided by VK_VERSION_1_1
typedef enum VkMemoryAllocateFlagBits {
    VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT = 0x00000001,
  // Provided by VK_VERSION_1_2
    VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT = 0x00000002,
  // Provided by VK_VERSION_1_2
    VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT = 0x00000004,
  // Provided by VK_KHR_device_group
    VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT_KHR = VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT,
  // Provided by VK_KHR_buffer_device_address
    VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT_KHR = VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT,
  // Provided by VK_KHR_buffer_device_address
    VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_KHR = VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT,
} VkMemoryAllocateFlagBits;
```

or the equivalent

``` c
// Provided by VK_KHR_device_group
typedef VkMemoryAllocateFlagBits VkMemoryAllocateFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT` specifies that memory will be
  allocated for the devices in
  [VkMemoryAllocateFlagsInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagsInfo.html)::`deviceMask`.

- `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT` specifies that the memory
  **can** be attached to a buffer object created with the
  `VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT` bit set in `usage`, and
  that the memory handle **can** be used to retrieve an opaque address
  via
  [vkGetDeviceMemoryOpaqueCaptureAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceMemoryOpaqueCaptureAddress.html).

- `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` specifies that
  the memory’s address **can** be saved and reused on a subsequent run
  (e.g. for trace capture and replay), see
  [VkBufferOpaqueCaptureAddressCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferOpaqueCaptureAddressCreateInfo.html)
  for more detail.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkMemoryAllocateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryAllocateFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
