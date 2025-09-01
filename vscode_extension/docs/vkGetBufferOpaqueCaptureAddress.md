# vkGetBufferOpaqueCaptureAddress(3) Manual Page

## Name

vkGetBufferOpaqueCaptureAddress - Query an opaque capture address of a buffer



## [](#_c_specification)C Specification

To query a 64-bit buffer opaque capture address, call:

```c++
// Provided by VK_VERSION_1_2
uint64_t vkGetBufferOpaqueCaptureAddress(
    VkDevice                                    device,
    const VkBufferDeviceAddressInfo*            pInfo);
```

or the equivalent command

```c++
// Provided by VK_KHR_buffer_device_address
uint64_t vkGetBufferOpaqueCaptureAddressKHR(
    VkDevice                                    device,
    const VkBufferDeviceAddressInfo*            pInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device that the buffer was created on.
- `pInfo` is a pointer to a [VkBufferDeviceAddressInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferDeviceAddressInfo.html) structure specifying the buffer to retrieve an address for.

## [](#_description)Description

The 64-bit return value is an opaque capture address of the start of `pInfo->buffer`.

If the buffer was created with a non-zero value of [VkBufferOpaqueCaptureAddressCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferOpaqueCaptureAddressCreateInfo.html)::`opaqueCaptureAddress` the return value **must** be the same address.

Valid Usage

- [](#VUID-vkGetBufferOpaqueCaptureAddress-None-03326)VUID-vkGetBufferOpaqueCaptureAddress-None-03326  
  The [`bufferDeviceAddress`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddress) and [`bufferDeviceAddressCaptureReplay`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressCaptureReplay) features **must** be enabled
- [](#VUID-vkGetBufferOpaqueCaptureAddress-pInfo-10725)VUID-vkGetBufferOpaqueCaptureAddress-pInfo-10725  
  `pInfo->buffer` **must** have been created with the `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` flag
- [](#VUID-vkGetBufferOpaqueCaptureAddress-device-03327)VUID-vkGetBufferOpaqueCaptureAddress-device-03327  
  If `device` was created with multiple physical devices, then the [`bufferDeviceAddressMultiDevice`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressMultiDevice) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkGetBufferOpaqueCaptureAddress-device-parameter)VUID-vkGetBufferOpaqueCaptureAddress-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetBufferOpaqueCaptureAddress-pInfo-parameter)VUID-vkGetBufferOpaqueCaptureAddress-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkBufferDeviceAddressInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferDeviceAddressInfo.html) structure

## [](#_see_also)See Also

[VK\_KHR\_buffer\_device\_address](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_buffer_device_address.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkBufferDeviceAddressInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferDeviceAddressInfo.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetBufferOpaqueCaptureAddress).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0