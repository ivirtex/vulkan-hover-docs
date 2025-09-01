# VkMemoryAllocateFlagBits(3) Manual Page

## Name

VkMemoryAllocateFlagBits - Bitmask specifying flags for a device memory allocation



## [](#_c_specification)C Specification

Bits which **can** be set in [VkMemoryAllocateFlagsInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateFlagsInfo.html)::`flags`, controlling device memory allocation, are:

```c++
// Provided by VK_VERSION_1_1
typedef enum VkMemoryAllocateFlagBits {
    VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT = 0x00000001,
  // Provided by VK_VERSION_1_2
    VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT = 0x00000002,
  // Provided by VK_VERSION_1_2
    VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT = 0x00000004,
  // Provided by VK_EXT_zero_initialize_device_memory
    VK_MEMORY_ALLOCATE_ZERO_INITIALIZE_BIT_EXT = 0x00000008,
  // Provided by VK_KHR_device_group
    VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT_KHR = VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT,
  // Provided by VK_KHR_buffer_device_address
    VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT_KHR = VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT,
  // Provided by VK_KHR_buffer_device_address
    VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_KHR = VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT,
} VkMemoryAllocateFlagBits;
```

or the equivalent

```c++
// Provided by VK_KHR_device_group
typedef VkMemoryAllocateFlagBits VkMemoryAllocateFlagBitsKHR;
```

## [](#_description)Description

- `VK_MEMORY_ALLOCATE_DEVICE_MASK_BIT` specifies that memory will be allocated for the devices in [VkMemoryAllocateFlagsInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateFlagsInfo.html)::`deviceMask`.
- `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT` specifies that the memory **can** be attached to a buffer object created with the `VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT` bit set in `usage`, and that the memory handle **can** be used to retrieve an opaque address via [vkGetDeviceMemoryOpaqueCaptureAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceMemoryOpaqueCaptureAddress.html).
- `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` specifies that the memory’s address **can** be saved and reused on a subsequent run (e.g. for trace capture and replay), see [VkBufferOpaqueCaptureAddressCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferOpaqueCaptureAddressCreateInfo.html) for more detail.
- `VK_MEMORY_ALLOCATE_ZERO_INITIALIZE_BIT_EXT` specifies that the memory will be zeroed automatically by the implementation before application is able to access it.

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkMemoryAllocateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryAllocateFlagBits).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0