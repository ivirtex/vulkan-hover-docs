# vkBindBufferMemory2(3) Manual Page

## Name

vkBindBufferMemory2 - Bind device memory to buffer objects



## <a href="#_c_specification" class="anchor"></a>C Specification

To attach memory to buffer objects for one or more buffers at a time,
call:

``` c
// Provided by VK_VERSION_1_1
VkResult vkBindBufferMemory2(
    VkDevice                                    device,
    uint32_t                                    bindInfoCount,
    const VkBindBufferMemoryInfo*               pBindInfos);
```

or the equivalent command

``` c
// Provided by VK_KHR_bind_memory2
VkResult vkBindBufferMemory2KHR(
    VkDevice                                    device,
    uint32_t                                    bindInfoCount,
    const VkBindBufferMemoryInfo*               pBindInfos);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the buffers and memory.

- `bindInfoCount` is the number of elements in `pBindInfos`.

- `pBindInfos` is a pointer to an array of `bindInfoCount`
  [VkBindBufferMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindBufferMemoryInfo.html) structures
  describing buffers and memory to bind.

## <a href="#_description" class="anchor"></a>Description

On some implementations, it **may** be more efficient to batch memory
bindings into a single command.

If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance6"
target="_blank" rel="noopener"><code>maintenance6</code></a> feature is
enabled, this command **must** attempt to perform all of the memory
binding operations described by `pBindInfos`, and **must** not early
exit on the first failure.

If any of the memory binding operations described by `pBindInfos` fail,
the [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html) returned by this command **must** be the
return value of any one of the memory binding operations which did not
return `VK_SUCCESS`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>If the <code>vkBindBufferMemory2</code> command failed, <a
href="VkBindMemoryStatusKHR.html">VkBindMemoryStatusKHR</a> structures
were not included in the <code>pNext</code> chains of each element of
<code>pBindInfos</code>, and <code>bindInfoCount</code> was greater than
one, then the buffers referenced by <code>pBindInfos</code> will be in
an indeterminate state, and must not be used.</p>
<p>Applications should destroy these buffers.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-vkBindBufferMemory2-device-parameter"
  id="VUID-vkBindBufferMemory2-device-parameter"></a>
  VUID-vkBindBufferMemory2-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkBindBufferMemory2-pBindInfos-parameter"
  id="VUID-vkBindBufferMemory2-pBindInfos-parameter"></a>
  VUID-vkBindBufferMemory2-pBindInfos-parameter  
  `pBindInfos` **must** be a valid pointer to an array of
  `bindInfoCount` valid
  [VkBindBufferMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindBufferMemoryInfo.html) structures

- <a href="#VUID-vkBindBufferMemory2-bindInfoCount-arraylength"
  id="VUID-vkBindBufferMemory2-bindInfoCount-arraylength"></a>
  VUID-vkBindBufferMemory2-bindInfoCount-arraylength  
  `bindInfoCount` **must** be greater than `0`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkBindBufferMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindBufferMemoryInfo.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkBindBufferMemory2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
