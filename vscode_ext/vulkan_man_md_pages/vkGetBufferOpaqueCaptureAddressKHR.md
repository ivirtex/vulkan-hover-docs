# vkGetBufferOpaqueCaptureAddress(3) Manual Page

## Name

vkGetBufferOpaqueCaptureAddress - Query an opaque capture address of a
buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To query a 64-bit buffer opaque capture address, call:

``` c
// Provided by VK_VERSION_1_2
uint64_t vkGetBufferOpaqueCaptureAddress(
    VkDevice                                    device,
    const VkBufferDeviceAddressInfo*            pInfo);
```

or the equivalent command

``` c
// Provided by VK_KHR_buffer_device_address
uint64_t vkGetBufferOpaqueCaptureAddressKHR(
    VkDevice                                    device,
    const VkBufferDeviceAddressInfo*            pInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that the buffer was created on.

- `pInfo` is a pointer to a
  [VkBufferDeviceAddressInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferDeviceAddressInfo.html) structure
  specifying the buffer to retrieve an address for.

## <a href="#_description" class="anchor"></a>Description

The 64-bit return value is an opaque capture address of the start of
`pInfo->buffer`.

If the buffer was created with a non-zero value of
[VkBufferOpaqueCaptureAddressCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferOpaqueCaptureAddressCreateInfo.html)::`opaqueCaptureAddress`
the return value **must** be the same address.

Valid Usage

- <a href="#VUID-vkGetBufferOpaqueCaptureAddress-None-03326"
  id="VUID-vkGetBufferOpaqueCaptureAddress-None-03326"></a>
  VUID-vkGetBufferOpaqueCaptureAddress-None-03326  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-bufferDeviceAddress"
  target="_blank" rel="noopener"><code>bufferDeviceAddress</code></a>
  feature **must** be enabled

- <a href="#VUID-vkGetBufferOpaqueCaptureAddress-device-03327"
  id="VUID-vkGetBufferOpaqueCaptureAddress-device-03327"></a>
  VUID-vkGetBufferOpaqueCaptureAddress-device-03327  
  If `device` was created with multiple physical devices, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-bufferDeviceAddressMultiDevice"
  target="_blank"
  rel="noopener"><code>bufferDeviceAddressMultiDevice</code></a> feature
  **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-vkGetBufferOpaqueCaptureAddress-device-parameter"
  id="VUID-vkGetBufferOpaqueCaptureAddress-device-parameter"></a>
  VUID-vkGetBufferOpaqueCaptureAddress-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetBufferOpaqueCaptureAddress-pInfo-parameter"
  id="VUID-vkGetBufferOpaqueCaptureAddress-pInfo-parameter"></a>
  VUID-vkGetBufferOpaqueCaptureAddress-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkBufferDeviceAddressInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferDeviceAddressInfo.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_buffer_device_address](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_buffer_device_address.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkBufferDeviceAddressInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferDeviceAddressInfo.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetBufferOpaqueCaptureAddress"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
