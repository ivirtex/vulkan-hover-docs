# vkGetMemoryRemoteAddressNV(3) Manual Page

## Name

vkGetMemoryRemoteAddressNV - Get an address for a memory object accessible by remote devices



## [](#_c_specification)C Specification

To export an address representing the payload of a Vulkan device memory object accessible by remote devices, call:

```c++
// Provided by VK_NV_external_memory_rdma
VkResult vkGetMemoryRemoteAddressNV(
    VkDevice                                    device,
    const VkMemoryGetRemoteAddressInfoNV*       pMemoryGetRemoteAddressInfo,
    VkRemoteAddressNV*                          pAddress);
```

## [](#_parameters)Parameters

- `device` is the logical device that created the device memory being exported.
- `pMemoryGetRemoteAddressInfo` is a pointer to a [VkMemoryGetRemoteAddressInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetRemoteAddressInfoNV.html) structure containing parameters of the export operation.
- `pAddress` is a pointer to a [VkRemoteAddressNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRemoteAddressNV.html) value in which an address representing the payload of the device memory object is returned.

## [](#_description)Description

More communication may be required between the kernel-mode drivers of the devices involved. This information is out of scope of this documentation and should be requested from the vendors of the devices.

Valid Usage (Implicit)

- [](#VUID-vkGetMemoryRemoteAddressNV-device-parameter)VUID-vkGetMemoryRemoteAddressNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetMemoryRemoteAddressNV-pMemoryGetRemoteAddressInfo-parameter)VUID-vkGetMemoryRemoteAddressNV-pMemoryGetRemoteAddressInfo-parameter  
  `pMemoryGetRemoteAddressInfo` **must** be a valid pointer to a valid [VkMemoryGetRemoteAddressInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetRemoteAddressInfoNV.html) structure
- [](#VUID-vkGetMemoryRemoteAddressNV-pAddress-parameter)VUID-vkGetMemoryRemoteAddressNV-pAddress-parameter  
  `pAddress` **must** be a valid pointer to a [VkRemoteAddressNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRemoteAddressNV.html) value

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INVALID_EXTERNAL_HANDLE`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_NV\_external\_memory\_rdma](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_external_memory_rdma.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkMemoryGetRemoteAddressInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetRemoteAddressInfoNV.html), [VkRemoteAddressNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRemoteAddressNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetMemoryRemoteAddressNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0