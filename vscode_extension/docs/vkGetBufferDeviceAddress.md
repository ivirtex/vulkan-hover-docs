# vkGetBufferDeviceAddress(3) Manual Page

## Name

vkGetBufferDeviceAddress - Query an address of a buffer



## [](#_c_specification)C Specification

To query a 64-bit buffer device address value which can be used to identify a buffer to API commands or through which buffer memory **can** be accessed, call:

```c++
// Provided by VK_VERSION_1_2
VkDeviceAddress vkGetBufferDeviceAddress(
    VkDevice                                    device,
    const VkBufferDeviceAddressInfo*            pInfo);
```

or the equivalent command

```c++
// Provided by VK_KHR_buffer_device_address
VkDeviceAddress vkGetBufferDeviceAddressKHR(
    VkDevice                                    device,
    const VkBufferDeviceAddressInfo*            pInfo);
```

or the equivalent command

```c++
// Provided by VK_EXT_buffer_device_address
VkDeviceAddress vkGetBufferDeviceAddressEXT(
    VkDevice                                    device,
    const VkBufferDeviceAddressInfo*            pInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device that the buffer was created on.
- `pInfo` is a pointer to a [VkBufferDeviceAddressInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferDeviceAddressInfo.html) structure specifying the buffer to retrieve an address for.

## [](#_description)Description

The 64-bit return value, `bufferBaseAddress`, is an address of the start of `pInfo->buffer`. Addresses in the range \[`bufferBaseAddress`, `bufferBaseAddress` + [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html)::`size`) **can** be used to access the memory bound to this buffer on the device.

A value of zero is reserved as a “null” pointer and **must** not be returned as a valid buffer device address.

If the buffer was created with a non-zero value of [VkBufferOpaqueCaptureAddressCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferOpaqueCaptureAddressCreateInfo.html)::`opaqueCaptureAddress` or [VkBufferDeviceAddressCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferDeviceAddressCreateInfoEXT.html)::`deviceAddress`, the return value will be the same address that was returned at capture time.

The returned address **must** satisfy the alignment requirement specified by [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html)::`alignment` for the buffer in [VkBufferDeviceAddressInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferDeviceAddressInfo.html)::`buffer`.

If multiple [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) objects are bound to overlapping ranges of [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), implementations **may** return address ranges which overlap. In this case, it is ambiguous which [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) is associated with any given device address. For purposes of valid usage, if multiple [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) objects **can** be attributed to a device address, a [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) is selected such that valid usage passes, if it exists.

Valid Usage

- [](#VUID-vkGetBufferDeviceAddress-bufferDeviceAddress-03324)VUID-vkGetBufferDeviceAddress-bufferDeviceAddress-03324  
  The [`bufferDeviceAddress`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddress) feature or the [`VkPhysicalDeviceBufferDeviceAddressFeaturesEXT`::`bufferDeviceAddress`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressEXT) feature **must** be enabled, and at least one of the following conditions **must** be met
  
  - `buffer` is sparse
  - `buffer` is bound completely and contiguously to a single `VkDeviceMemory` object
  - `buffer` was created with the `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` flag and the [`VkPhysicalDeviceBufferDeviceAddressFeaturesEXT`::`bufferDeviceAddress`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressEXT) feature is enabled on the device
- [](#VUID-vkGetBufferDeviceAddress-device-03325)VUID-vkGetBufferDeviceAddress-device-03325  
  If `device` was created with multiple physical devices, then the [`bufferDeviceAddressMultiDevice`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressMultiDevice) or [`VkPhysicalDeviceBufferDeviceAddressFeaturesEXT`::`bufferDeviceAddressMultiDevice`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-bufferDeviceAddressMultiDeviceEXT) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkGetBufferDeviceAddress-device-parameter)VUID-vkGetBufferDeviceAddress-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetBufferDeviceAddress-pInfo-parameter)VUID-vkGetBufferDeviceAddress-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkBufferDeviceAddressInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferDeviceAddressInfo.html) structure

## [](#_see_also)See Also

[VK\_KHR\_buffer\_device\_address](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_buffer_device_address.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkBufferDeviceAddressInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferDeviceAddressInfo.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetBufferDeviceAddress)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0