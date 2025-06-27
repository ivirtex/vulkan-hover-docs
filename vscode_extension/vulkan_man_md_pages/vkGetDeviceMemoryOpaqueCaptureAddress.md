# vkGetDeviceMemoryOpaqueCaptureAddress(3) Manual Page

## Name

vkGetDeviceMemoryOpaqueCaptureAddress - Query an opaque capture address
of a memory object



## <a href="#_c_specification" class="anchor"></a>C Specification

To query a 64-bit opaque capture address value from a memory object,
call:

``` c
// Provided by VK_VERSION_1_2
uint64_t vkGetDeviceMemoryOpaqueCaptureAddress(
    VkDevice                                    device,
    const VkDeviceMemoryOpaqueCaptureAddressInfo* pInfo);
```

or the equivalent command

``` c
// Provided by VK_KHR_buffer_device_address
uint64_t vkGetDeviceMemoryOpaqueCaptureAddressKHR(
    VkDevice                                    device,
    const VkDeviceMemoryOpaqueCaptureAddressInfo* pInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that the memory object was allocated
  on.

- `pInfo` is a pointer to a
  [VkDeviceMemoryOpaqueCaptureAddressInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryOpaqueCaptureAddressInfo.html)
  structure specifying the memory object to retrieve an address for.

## <a href="#_description" class="anchor"></a>Description

The 64-bit return value is an opaque address representing the start of
`pInfo->memory`.

If the memory object was allocated with a non-zero value of
[VkMemoryOpaqueCaptureAddressAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryOpaqueCaptureAddressAllocateInfo.html)::`opaqueCaptureAddress`,
the return value **must** be the same address.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The expected usage for these opaque addresses is only for trace
capture/replay tools to store these addresses in a trace and
subsequently specify them during replay.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkGetDeviceMemoryOpaqueCaptureAddress-None-03334"
  id="VUID-vkGetDeviceMemoryOpaqueCaptureAddress-None-03334"></a>
  VUID-vkGetDeviceMemoryOpaqueCaptureAddress-None-03334  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-bufferDeviceAddress"
  target="_blank" rel="noopener"><code>bufferDeviceAddress</code></a>
  feature **must** be enabled

- <a href="#VUID-vkGetDeviceMemoryOpaqueCaptureAddress-device-03335"
  id="VUID-vkGetDeviceMemoryOpaqueCaptureAddress-device-03335"></a>
  VUID-vkGetDeviceMemoryOpaqueCaptureAddress-device-03335  
  If `device` was created with multiple physical devices, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-bufferDeviceAddressMultiDevice"
  target="_blank"
  rel="noopener"><code>bufferDeviceAddressMultiDevice</code></a> feature
  **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-vkGetDeviceMemoryOpaqueCaptureAddress-device-parameter"
  id="VUID-vkGetDeviceMemoryOpaqueCaptureAddress-device-parameter"></a>
  VUID-vkGetDeviceMemoryOpaqueCaptureAddress-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetDeviceMemoryOpaqueCaptureAddress-pInfo-parameter"
  id="VUID-vkGetDeviceMemoryOpaqueCaptureAddress-pInfo-parameter"></a>
  VUID-vkGetDeviceMemoryOpaqueCaptureAddress-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkDeviceMemoryOpaqueCaptureAddressInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryOpaqueCaptureAddressInfo.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_buffer_device_address](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_buffer_device_address.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkDeviceMemoryOpaqueCaptureAddressInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryOpaqueCaptureAddressInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDeviceMemoryOpaqueCaptureAddress"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
