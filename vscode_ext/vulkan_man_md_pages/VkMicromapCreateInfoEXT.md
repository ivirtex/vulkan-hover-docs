# VkMicromapCreateInfoEXT(3) Manual Page

## Name

VkMicromapCreateInfoEXT - Structure specifying the parameters of a newly
created micromap object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMicromapCreateInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_opacity_micromap
typedef struct VkMicromapCreateInfoEXT {
    VkStructureType             sType;
    const void*                 pNext;
    VkMicromapCreateFlagsEXT    createFlags;
    VkBuffer                    buffer;
    VkDeviceSize                offset;
    VkDeviceSize                size;
    VkMicromapTypeEXT           type;
    VkDeviceAddress             deviceAddress;
} VkMicromapCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `createFlags` is a bitmask of
  [VkMicromapCreateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapCreateFlagBitsEXT.html)
  specifying additional creation parameters of the micromap.

- `buffer` is the buffer on which the micromap will be stored.

- `offset` is an offset in bytes from the base address of the buffer at
  which the micromap will be stored, and **must** be a multiple of
  `256`.

- `size` is the size required for the micromap.

- `type` is a [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTypeEXT.html) value
  specifying the type of micromap that will be created.

- `deviceAddress` is the device address requested for the micromap if
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-micromapCaptureReplay"
  target="_blank" rel="noopener"><code>micromapCaptureReplay</code></a>
  feature is being used.

## <a href="#_description" class="anchor"></a>Description

If `deviceAddress` is zero, no specific address is requested.

If `deviceAddress` is not zero, `deviceAddress` **must** be an address
retrieved from an identically created micromap on the same
implementation. The micromap **must** also be placed on an identically
created `buffer` and at the same `offset`.

Applications **should** avoid creating micromaps with
application-provided addresses and implementation-provided addresses in
the same process, to reduce the likelihood of
`VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR` errors.

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
<code>VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT</code>, and will add
<code>VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT</code> to all buffers
used as storage for a micromap where <code>deviceAddress</code> is not
zero. This also means that the tool will need to add
<code>VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT</code> to memory allocations
to allow the flag to be set where the application may not have otherwise
required it. During capture the tool will save the queried opaque device
addresses in the trace. During replay, the buffers will be created
specifying the original address so any address values stored in the
trace data will remain valid.</p>
<p>Implementations are expected to separate such buffers in the GPU
address space so normal allocations will avoid using these addresses.
Apps/tools should avoid mixing app-provided and implementation-provided
addresses for buffers created with
<code>VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT</code>, to
avoid address space allocation conflicts.</p></td>
</tr>
</tbody>
</table>

If the micromap will be the target of a build operation, the required
size for a micromap **can** be queried with
[vkGetMicromapBuildSizesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMicromapBuildSizesEXT.html).

Valid Usage

- <a href="#VUID-VkMicromapCreateInfoEXT-deviceAddress-07433"
  id="VUID-VkMicromapCreateInfoEXT-deviceAddress-07433"></a>
  VUID-VkMicromapCreateInfoEXT-deviceAddress-07433  
  If `deviceAddress` is not zero, `createFlags` **must** include
  `VK_MICROMAP_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_EXT`

- <a href="#VUID-VkMicromapCreateInfoEXT-createFlags-07434"
  id="VUID-VkMicromapCreateInfoEXT-createFlags-07434"></a>
  VUID-VkMicromapCreateInfoEXT-createFlags-07434  
  If `createFlags` includes
  `VK_MICROMAP_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_EXT`,
  [VkPhysicalDeviceOpacityMicromapFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceOpacityMicromapFeaturesEXT.html)::`micromapCaptureReplay`
  **must** be `VK_TRUE`

- <a href="#VUID-VkMicromapCreateInfoEXT-buffer-07435"
  id="VUID-VkMicromapCreateInfoEXT-buffer-07435"></a>
  VUID-VkMicromapCreateInfoEXT-buffer-07435  
  `buffer` **must** have been created with a `usage` value containing
  `VK_BUFFER_USAGE_MICROMAP_STORAGE_BIT_EXT`

- <a href="#VUID-VkMicromapCreateInfoEXT-buffer-07436"
  id="VUID-VkMicromapCreateInfoEXT-buffer-07436"></a>
  VUID-VkMicromapCreateInfoEXT-buffer-07436  
  `buffer` **must** not have been created with
  `VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT`

- <a href="#VUID-VkMicromapCreateInfoEXT-offset-07437"
  id="VUID-VkMicromapCreateInfoEXT-offset-07437"></a>
  VUID-VkMicromapCreateInfoEXT-offset-07437  
  The sum of `offset` and `size` **must** be less than the size of
  `buffer`

- <a href="#VUID-VkMicromapCreateInfoEXT-offset-07438"
  id="VUID-VkMicromapCreateInfoEXT-offset-07438"></a>
  VUID-VkMicromapCreateInfoEXT-offset-07438  
  `offset` **must** be a multiple of `256` bytes

Valid Usage (Implicit)

- <a href="#VUID-VkMicromapCreateInfoEXT-sType-sType"
  id="VUID-VkMicromapCreateInfoEXT-sType-sType"></a>
  VUID-VkMicromapCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MICROMAP_CREATE_INFO_EXT`

- <a href="#VUID-VkMicromapCreateInfoEXT-pNext-pNext"
  id="VUID-VkMicromapCreateInfoEXT-pNext-pNext"></a>
  VUID-VkMicromapCreateInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkMicromapCreateInfoEXT-createFlags-parameter"
  id="VUID-VkMicromapCreateInfoEXT-createFlags-parameter"></a>
  VUID-VkMicromapCreateInfoEXT-createFlags-parameter  
  `createFlags` **must** be a valid combination of
  [VkMicromapCreateFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapCreateFlagBitsEXT.html) values

- <a href="#VUID-VkMicromapCreateInfoEXT-buffer-parameter"
  id="VUID-VkMicromapCreateInfoEXT-buffer-parameter"></a>
  VUID-VkMicromapCreateInfoEXT-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-VkMicromapCreateInfoEXT-type-parameter"
  id="VUID-VkMicromapCreateInfoEXT-type-parameter"></a>
  VUID-VkMicromapCreateInfoEXT-type-parameter  
  `type` **must** be a valid [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTypeEXT.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkMicromapCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapCreateFlagsEXT.html),
[VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTypeEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateMicromapEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMicromapCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
