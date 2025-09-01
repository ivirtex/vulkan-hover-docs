# vkGetDeviceMemoryOpaqueCaptureAddress(3) Manual Page

## Name

vkGetDeviceMemoryOpaqueCaptureAddress - Query an opaque capture address of a memory object



## [](#_c_specification)C Specification

To query a 64-bit opaque capture address value from a memory object, call:

```c++
// Provided by VK_VERSION_1_2
uint64_t vkGetDeviceMemoryOpaqueCaptureAddress(
    VkDevice                                    device,
    const VkDeviceMemoryOpaqueCaptureAddressInfo* pInfo);
```

or the equivalent command

```c++
// Provided by VK_KHR_buffer_device_address
uint64_t vkGetDeviceMemoryOpaqueCaptureAddressKHR(
    VkDevice                                    device,
    const VkDeviceMemoryOpaqueCaptureAddressInfo* pInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device that the memory object was allocated on.
- `pInfo` is a pointer to a [VkDeviceMemoryOpaqueCaptureAddressInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryOpaqueCaptureAddressInfo.html) structure specifying the memory object to retrieve an address for.

## [](#_description)Description

The 64-bit return value is an opaque address representing the start of `pInfo->memory`.

If the memory object was allocated with a non-zero value of [VkMemoryOpaqueCaptureAddressAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryOpaqueCaptureAddressAllocateInfo.html)::`opaqueCaptureAddress`, the return value **must** be the same address.

Note

The expected usage for these opaque addresses is only for trace capture/replay tools to store these addresses in a trace and subsequently specify them during replay.

Valid Usage

- [](#VUID-vkGetDeviceMemoryOpaqueCaptureAddress-None-03334)VUID-vkGetDeviceMemoryOpaqueCaptureAddress-None-03334  
  The [`bufferDeviceAddress`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddress) and [`bufferDeviceAddressCaptureReplay`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressCaptureReplay) features **must** be enabled
- [](#VUID-vkGetDeviceMemoryOpaqueCaptureAddress-pInfo-10727)VUID-vkGetDeviceMemoryOpaqueCaptureAddress-pInfo-10727  
  `pInfo->memory` **must** have been allocated using the `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` flag
- [](#VUID-vkGetDeviceMemoryOpaqueCaptureAddress-device-03335)VUID-vkGetDeviceMemoryOpaqueCaptureAddress-device-03335  
  If `device` was created with multiple physical devices, then the [`bufferDeviceAddressMultiDevice`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressMultiDevice) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkGetDeviceMemoryOpaqueCaptureAddress-device-parameter)VUID-vkGetDeviceMemoryOpaqueCaptureAddress-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDeviceMemoryOpaqueCaptureAddress-pInfo-parameter)VUID-vkGetDeviceMemoryOpaqueCaptureAddress-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkDeviceMemoryOpaqueCaptureAddressInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryOpaqueCaptureAddressInfo.html) structure

## [](#_see_also)See Also

[VK\_KHR\_buffer\_device\_address](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_buffer_device_address.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkDeviceMemoryOpaqueCaptureAddressInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryOpaqueCaptureAddressInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDeviceMemoryOpaqueCaptureAddress).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0