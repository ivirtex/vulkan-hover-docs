# VkBufferOpaqueCaptureAddressCreateInfo(3) Manual Page

## Name

VkBufferOpaqueCaptureAddressCreateInfo - Request a specific address for
a buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To request a specific device address for a buffer, add a
[VkBufferOpaqueCaptureAddressCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferOpaqueCaptureAddressCreateInfo.html)
structure to the `pNext` chain of the
[VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html) structure. The
`VkBufferOpaqueCaptureAddressCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkBufferOpaqueCaptureAddressCreateInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint64_t           opaqueCaptureAddress;
} VkBufferOpaqueCaptureAddressCreateInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_buffer_device_address
typedef VkBufferOpaqueCaptureAddressCreateInfo VkBufferOpaqueCaptureAddressCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `opaqueCaptureAddress` is the opaque capture address requested for the
  buffer.

## <a href="#_description" class="anchor"></a>Description

If `opaqueCaptureAddress` is zero, no specific address is requested.

If `opaqueCaptureAddress` is not zero, then it **should** be an address
retrieved from
[vkGetBufferOpaqueCaptureAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferOpaqueCaptureAddress.html)
for an identically created buffer on the same implementation.

If this structure is not present, it is as if `opaqueCaptureAddress` is
zero.

Apps **should** avoid creating buffers with app-provided addresses and
implementation-provided addresses in the same process, to reduce the
likelihood of `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS` errors.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The expected usage for this is that a trace capture/replay tool will
add the <code>VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT</code>
flag to all buffers that use
<code>VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT</code>, and during
capture will save the queried opaque device addresses in the trace.
During replay, the buffers will be created specifying the original
address so any address values stored in the trace data will remain
valid.</p>
<p>Implementations are expected to separate such buffers in the GPU
address space so normal allocations will avoid using these addresses.
Apps/tools should avoid mixing app-provided and implementation-provided
addresses for buffers created with
<code>VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT</code>, to
avoid address space allocation conflicts.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkBufferOpaqueCaptureAddressCreateInfo-sType-sType"
  id="VUID-VkBufferOpaqueCaptureAddressCreateInfo-sType-sType"></a>
  VUID-VkBufferOpaqueCaptureAddressCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_BUFFER_OPAQUE_CAPTURE_ADDRESS_CREATE_INFO`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_buffer_device_address](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_buffer_device_address.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBufferOpaqueCaptureAddressCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
