# vkGetMemoryRemoteAddressNV(3) Manual Page

## Name

vkGetMemoryRemoteAddressNV - Get an address for a memory object
accessible by remote devices



## <a href="#_c_specification" class="anchor"></a>C Specification

To export an address representing the payload of a Vulkan device memory
object accessible by remote devices, call:

``` c
// Provided by VK_NV_external_memory_rdma
VkResult vkGetMemoryRemoteAddressNV(
    VkDevice                                    device,
    const VkMemoryGetRemoteAddressInfoNV*       pMemoryGetRemoteAddressInfo,
    VkRemoteAddressNV*                          pAddress);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that created the device memory being
  exported.

- `pMemoryGetRemoteAddressInfo` is a pointer to a
  [VkMemoryGetRemoteAddressInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryGetRemoteAddressInfoNV.html)
  structure containing parameters of the export operation.

- `pAddress` is a pointer to a
  [VkRemoteAddressNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRemoteAddressNV.html) value in which an address
  representing the payload of the device memory object is returned.

## <a href="#_description" class="anchor"></a>Description

More communication may be required between the kernel-mode drivers of
the devices involved. This information is out of scope of this
documentation and should be requested from the vendors of the devices.

Valid Usage (Implicit)

- <a href="#VUID-vkGetMemoryRemoteAddressNV-device-parameter"
  id="VUID-vkGetMemoryRemoteAddressNV-device-parameter"></a>
  VUID-vkGetMemoryRemoteAddressNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetMemoryRemoteAddressNV-pMemoryGetRemoteAddressInfo-parameter"
  id="VUID-vkGetMemoryRemoteAddressNV-pMemoryGetRemoteAddressInfo-parameter"></a>
  VUID-vkGetMemoryRemoteAddressNV-pMemoryGetRemoteAddressInfo-parameter  
  `pMemoryGetRemoteAddressInfo` **must** be a valid pointer to a valid
  [VkMemoryGetRemoteAddressInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryGetRemoteAddressInfoNV.html)
  structure

- <a href="#VUID-vkGetMemoryRemoteAddressNV-pAddress-parameter"
  id="VUID-vkGetMemoryRemoteAddressNV-pAddress-parameter"></a>
  VUID-vkGetMemoryRemoteAddressNV-pAddress-parameter  
  `pAddress` **must** be a valid pointer to a
  [VkRemoteAddressNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRemoteAddressNV.html) value

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_INVALID_EXTERNAL_HANDLE`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_external_memory_rdma](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_external_memory_rdma.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkMemoryGetRemoteAddressInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryGetRemoteAddressInfoNV.html),
[VkRemoteAddressNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRemoteAddressNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetMemoryRemoteAddressNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
