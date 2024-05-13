# vkCreateMicromapEXT(3) Manual Page

## Name

vkCreateMicromapEXT - Create a new micromap object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a micromap, call:

``` c
// Provided by VK_EXT_opacity_micromap
VkResult vkCreateMicromapEXT(
    VkDevice                                    device,
    const VkMicromapCreateInfoEXT*              pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkMicromapEXT*                              pMicromap);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the acceleration structure
  object.

- `pCreateInfo` is a pointer to a
  [VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapCreateInfoEXT.html) structure
  containing parameters affecting creation of the micromap.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pMicromap` is a pointer to a `VkMicromapEXT` handle in which the
  resulting micromap object is returned.

## <a href="#_description" class="anchor"></a>Description

Similar to other objects in Vulkan, the micromap creation merely creates
an object with a specific “shape”. The type and quantity of geometry
that can be built into a micromap is determined by the parameters of
[VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapCreateInfoEXT.html).

The micromap data is stored in the object referred to by
`VkMicromapCreateInfoEXT`::`buffer`. Once memory has been bound to that
buffer, it **must** be populated by micromap build or micromap copy
commands such as [vkCmdBuildMicromapsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBuildMicromapsEXT.html),
[vkBuildMicromapsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBuildMicromapsEXT.html),
[vkCmdCopyMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMicromapEXT.html), and
[vkCopyMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyMicromapEXT.html).

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The expected usage for a trace capture/replay tool is that it will
serialize and later deserialize the micromap data using micromap copy
commands. During capture the tool will use <a
href="vkCopyMicromapToMemoryEXT.html">vkCopyMicromapToMemoryEXT</a> or
<a
href="vkCmdCopyMicromapToMemoryEXT.html">vkCmdCopyMicromapToMemoryEXT</a>
with a <code>mode</code> of
<code>VK_COPY_MICROMAP_MODE_SERIALIZE_EXT</code>, and <a
href="vkCopyMemoryToMicromapEXT.html">vkCopyMemoryToMicromapEXT</a> or
<a
href="vkCmdCopyMemoryToMicromapEXT.html">vkCmdCopyMemoryToMicromapEXT</a>
with a <code>mode</code> of
<code>VK_COPY_MICROMAP_MODE_DESERIALIZE_EXT</code> during
replay.</p></td>
</tr>
</tbody>
</table>

The input buffers passed to micromap build commands will be referenced
by the implementation for the duration of the command. Micromaps
**must** be fully self-contained. The application **can** reuse or free
any memory which was used by the command as an input or as scratch
without affecting the results of a subsequent acceleration structure
build using the micromap or traversal of that acceleration structure.

Valid Usage

- <a href="#VUID-vkCreateMicromapEXT-micromap-07430"
  id="VUID-vkCreateMicromapEXT-micromap-07430"></a>
  VUID-vkCreateMicromapEXT-micromap-07430  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-micromap"
  target="_blank" rel="noopener"><code>micromap</code></a> feature
  **must** be enabled

- <a href="#VUID-vkCreateMicromapEXT-deviceAddress-07431"
  id="VUID-vkCreateMicromapEXT-deviceAddress-07431"></a>
  VUID-vkCreateMicromapEXT-deviceAddress-07431  
  If
  [VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapCreateInfoEXT.html)::`deviceAddress`
  is not zero, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-micromapCaptureReplay"
  target="_blank" rel="noopener"><code>micromapCaptureReplay</code></a>
  feature **must** be enabled

- <a href="#VUID-vkCreateMicromapEXT-device-07432"
  id="VUID-vkCreateMicromapEXT-device-07432"></a>
  VUID-vkCreateMicromapEXT-device-07432  
  If `device` was created with multiple physical devices, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-bufferDeviceAddressMultiDevice"
  target="_blank"
  rel="noopener"><code>bufferDeviceAddressMultiDevice</code></a> feature
  **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-vkCreateMicromapEXT-device-parameter"
  id="VUID-vkCreateMicromapEXT-device-parameter"></a>
  VUID-vkCreateMicromapEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateMicromapEXT-pCreateInfo-parameter"
  id="VUID-vkCreateMicromapEXT-pCreateInfo-parameter"></a>
  VUID-vkCreateMicromapEXT-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapCreateInfoEXT.html) structure

- <a href="#VUID-vkCreateMicromapEXT-pAllocator-parameter"
  id="VUID-vkCreateMicromapEXT-pAllocator-parameter"></a>
  VUID-vkCreateMicromapEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateMicromapEXT-pMicromap-parameter"
  id="VUID-vkCreateMicromapEXT-pMicromap-parameter"></a>
  VUID-vkCreateMicromapEXT-pMicromap-parameter  
  `pMicromap` **must** be a valid pointer to a
  [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapCreateInfoEXT.html),
[VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateMicromapEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
