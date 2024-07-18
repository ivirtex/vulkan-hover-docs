# vkBindImageMemory2(3) Manual Page

## Name

vkBindImageMemory2 - Bind device memory to image objects



## <a href="#_c_specification" class="anchor"></a>C Specification

To attach memory to image objects for one or more images at a time,
call:

``` c
// Provided by VK_VERSION_1_1
VkResult vkBindImageMemory2(
    VkDevice                                    device,
    uint32_t                                    bindInfoCount,
    const VkBindImageMemoryInfo*                pBindInfos);
```

or the equivalent command

``` c
// Provided by VK_KHR_bind_memory2
VkResult vkBindImageMemory2KHR(
    VkDevice                                    device,
    uint32_t                                    bindInfoCount,
    const VkBindImageMemoryInfo*                pBindInfos);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the images and memory.

- `bindInfoCount` is the number of elements in `pBindInfos`.

- `pBindInfos` is a pointer to an array of
  [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryInfo.html) structures,
  describing images and memory to bind.

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
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>If the <code>vkBindImageMemory2</code> command failed, <a
href="VkBindMemoryStatusKHR.html">VkBindMemoryStatusKHR</a> structures
were not included in the <code>pNext</code> chains of each element of
<code>pBindInfos</code>, and <code>bindInfoCount</code> was greater than
one, then the images referenced by <code>pBindInfos</code> will be in an
indeterminate state, and must not be used.</p>
<p>Applications should destroy these images.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkBindImageMemory2-pBindInfos-02858"
  id="VUID-vkBindImageMemory2-pBindInfos-02858"></a>
  VUID-vkBindImageMemory2-pBindInfos-02858  
  If any [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryInfo.html)::`image`
  was created with `VK_IMAGE_CREATE_DISJOINT_BIT` then all planes of
  [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryInfo.html)::`image` **must**
  be bound individually in separate `pBindInfos`

- <a href="#VUID-vkBindImageMemory2-pBindInfos-04006"
  id="VUID-vkBindImageMemory2-pBindInfos-04006"></a>
  VUID-vkBindImageMemory2-pBindInfos-04006  
  `pBindInfos` **must** not refer to the same image subresource more
  than once

Valid Usage (Implicit)

- <a href="#VUID-vkBindImageMemory2-device-parameter"
  id="VUID-vkBindImageMemory2-device-parameter"></a>
  VUID-vkBindImageMemory2-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkBindImageMemory2-pBindInfos-parameter"
  id="VUID-vkBindImageMemory2-pBindInfos-parameter"></a>
  VUID-vkBindImageMemory2-pBindInfos-parameter  
  `pBindInfos` **must** be a valid pointer to an array of
  `bindInfoCount` valid
  [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryInfo.html) structures

- <a href="#VUID-vkBindImageMemory2-bindInfoCount-arraylength"
  id="VUID-vkBindImageMemory2-bindInfoCount-arraylength"></a>
  VUID-vkBindImageMemory2-bindInfoCount-arraylength  
  `bindInfoCount` **must** be greater than `0`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryInfo.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkBindImageMemory2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
