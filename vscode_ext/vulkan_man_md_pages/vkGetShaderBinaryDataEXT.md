# vkGetShaderBinaryDataEXT(3) Manual Page

## Name

vkGetShaderBinaryDataEXT - Get the binary shader code from a shader
object



## <a href="#_c_specification" class="anchor"></a>C Specification

Binary shader code **can** be retrieved from a shader object using the
command:

``` c
// Provided by VK_EXT_shader_object
VkResult vkGetShaderBinaryDataEXT(
    VkDevice                                    device,
    VkShaderEXT                                 shader,
    size_t*                                     pDataSize,
    void*                                       pData);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that shader object was created from.

- `shader` is the shader object to retrieve binary shader code from.

- `pDataSize` is a pointer to a `size_t` value related to the size of
  the binary shader code, as described below.

- `pData` is either `NULL` or a pointer to a buffer.

## <a href="#_description" class="anchor"></a>Description

If `pData` is `NULL`, then the size of the binary shader code of the
shader object, in bytes, is returned in `pDataSize`. Otherwise,
`pDataSize` **must** point to a variable set by the application to the
size of the buffer, in bytes, pointed to by `pData`, and on return the
variable is overwritten with the amount of data actually written to
`pData`. If `pDataSize` is less than the size of the binary shader code,
nothing is written to `pData`, and `VK_INCOMPLETE` will be returned
instead of `VK_SUCCESS`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The behavior of this command when <code>pDataSize</code> is too small
differs from how some other getter-type commands work in Vulkan. Because
shader binary data is only usable in its entirety, it would never be
useful for the implementation to return partial data. Because of this,
nothing is written to <code>pData</code> unless <code>pDataSize</code>
is large enough to fit the data in its entirety.</p></td>
</tr>
</tbody>
</table>

Binary shader code retrieved using `vkGetShaderBinaryDataEXT` **can** be
passed to a subsequent call to
[vkCreateShadersEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateShadersEXT.html) on a compatible physical
device by specifying `VK_SHADER_CODE_TYPE_BINARY_EXT` in the `codeType`
member of `VkShaderCreateInfoEXT`.

The shader code returned by repeated calls to this function with the
same `VkShaderEXT` is guaranteed to be invariant for the lifetime of the
`VkShaderEXT` object.

Valid Usage

- <a href="#VUID-vkGetShaderBinaryDataEXT-None-08461"
  id="VUID-vkGetShaderBinaryDataEXT-None-08461"></a>
  VUID-vkGetShaderBinaryDataEXT-None-08461  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderObject"
  target="_blank" rel="noopener"><code>shaderObject</code></a> feature
  **must** be enabled

- <a href="#VUID-vkGetShaderBinaryDataEXT-None-08499"
  id="VUID-vkGetShaderBinaryDataEXT-None-08499"></a>
  VUID-vkGetShaderBinaryDataEXT-None-08499  
  If `pData` is not `NULL`, it **must** be aligned to `16` bytes

Valid Usage (Implicit)

- <a href="#VUID-vkGetShaderBinaryDataEXT-device-parameter"
  id="VUID-vkGetShaderBinaryDataEXT-device-parameter"></a>
  VUID-vkGetShaderBinaryDataEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetShaderBinaryDataEXT-shader-parameter"
  id="VUID-vkGetShaderBinaryDataEXT-shader-parameter"></a>
  VUID-vkGetShaderBinaryDataEXT-shader-parameter  
  `shader` **must** be a valid [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html) handle

- <a href="#VUID-vkGetShaderBinaryDataEXT-pDataSize-parameter"
  id="VUID-vkGetShaderBinaryDataEXT-pDataSize-parameter"></a>
  VUID-vkGetShaderBinaryDataEXT-pDataSize-parameter  
  `pDataSize` **must** be a valid pointer to a `size_t` value

- <a href="#VUID-vkGetShaderBinaryDataEXT-pData-parameter"
  id="VUID-vkGetShaderBinaryDataEXT-pData-parameter"></a>
  VUID-vkGetShaderBinaryDataEXT-pData-parameter  
  If the value referenced by `pDataSize` is not `0`, and `pData` is not
  `NULL`, `pData` **must** be a valid pointer to an array of `pDataSize`
  bytes

- <a href="#VUID-vkGetShaderBinaryDataEXT-shader-parent"
  id="VUID-vkGetShaderBinaryDataEXT-shader-parent"></a>
  VUID-vkGetShaderBinaryDataEXT-shader-parent  
  `shader` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetShaderBinaryDataEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
