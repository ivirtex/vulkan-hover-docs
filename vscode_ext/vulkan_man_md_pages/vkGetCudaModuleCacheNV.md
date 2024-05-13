# vkGetCudaModuleCacheNV(3) Manual Page

## Name

vkGetCudaModuleCacheNV - Get CUDA module cache



## <a href="#_c_specification" class="anchor"></a>C Specification

To get the CUDA module cache call:

``` c
// Provided by VK_NV_cuda_kernel_launch
VkResult vkGetCudaModuleCacheNV(
    VkDevice                                    device,
    VkCudaModuleNV                              module,
    size_t*                                     pCacheSize,
    void*                                       pCacheData);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the Function.

- `module` is the CUDA module.

- `pCacheSize` is a pointer containing the amount of bytes to be copied
  in `pCacheData`

- `pCacheData` is a pointer to a buffer in which to copy the binary
  cache

## <a href="#_description" class="anchor"></a>Description

If `pCacheData` is `NULL`, then the size of the binary cache, in bytes,
is returned in `pCacheSize`. Otherwise, `pCacheSize` **must** point to a
variable set by the user to the size of the buffer, in bytes, pointed to
by `pCacheData`, and on return the variable is overwritten with the
amount of data actually written to `pCacheData`. If `pCacheSize` is less
than the size of the binary shader code, nothing is written to
`pCacheData`, and `VK_INCOMPLETE` will be returned instead of
`VK_SUCCESS`.

The returned cache **may** then be used later for further initialization
of the CUDA module, by sending this cache *instead* of the PTX code when
using [vkCreateCudaModuleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateCudaModuleNV.html).

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Using the binary cache instead of the original PTX code
<strong>should</strong> significantly speed up initialization of the
CUDA module, given that the whole compilation and validation will not be
necessary.</p>
<p>As with <a href="VkPipelineCache.html">VkPipelineCache</a>, the
binary cache depends on the specific implementation. The application
<strong>must</strong> assume the cache upload might fail in many
circumstances and thus <strong>may</strong> have to get ready for
falling back to the original PTX code if necessary. Most often, the
cache <strong>may</strong> succeed if the same device driver and
architecture is used between the cache generation from PTX and the use
of this cache. In the event of a new driver version or if using a
different device architecture, this cache <strong>may</strong> become
invalid.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-vkGetCudaModuleCacheNV-device-parameter"
  id="VUID-vkGetCudaModuleCacheNV-device-parameter"></a>
  VUID-vkGetCudaModuleCacheNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetCudaModuleCacheNV-module-parameter"
  id="VUID-vkGetCudaModuleCacheNV-module-parameter"></a>
  VUID-vkGetCudaModuleCacheNV-module-parameter  
  `module` **must** be a valid [VkCudaModuleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaModuleNV.html)
  handle

- <a href="#VUID-vkGetCudaModuleCacheNV-pCacheSize-parameter"
  id="VUID-vkGetCudaModuleCacheNV-pCacheSize-parameter"></a>
  VUID-vkGetCudaModuleCacheNV-pCacheSize-parameter  
  `pCacheSize` **must** be a valid pointer to a `size_t` value

- <a href="#VUID-vkGetCudaModuleCacheNV-pCacheData-parameter"
  id="VUID-vkGetCudaModuleCacheNV-pCacheData-parameter"></a>
  VUID-vkGetCudaModuleCacheNV-pCacheData-parameter  
  If the value referenced by `pCacheSize` is not `0`, and `pCacheData`
  is not `NULL`, `pCacheData` **must** be a valid pointer to an array of
  `pCacheSize` bytes

- <a href="#VUID-vkGetCudaModuleCacheNV-module-parent"
  id="VUID-vkGetCudaModuleCacheNV-module-parent"></a>
  VUID-vkGetCudaModuleCacheNV-module-parent  
  `module` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_INITIALIZATION_FAILED`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_cuda_kernel_launch](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_cuda_kernel_launch.html),
[VkCudaModuleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCudaModuleNV.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetCudaModuleCacheNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
