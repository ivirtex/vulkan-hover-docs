# vkCreateValidationCacheEXT(3) Manual Page

## Name

vkCreateValidationCacheEXT - Creates a new validation cache



## <a href="#_c_specification" class="anchor"></a>C Specification

To create validation cache objects, call:

``` c
// Provided by VK_EXT_validation_cache
VkResult vkCreateValidationCacheEXT(
    VkDevice                                    device,
    const VkValidationCacheCreateInfoEXT*       pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkValidationCacheEXT*                       pValidationCache);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the validation cache
  object.

- `pCreateInfo` is a pointer to a
  [VkValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheCreateInfoEXT.html)
  structure containing the initial parameters for the validation cache
  object.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pValidationCache` is a pointer to a
  [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheEXT.html) handle in which the
  resulting validation cache object is returned.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Applications <strong>can</strong> track and manage the total host
memory size of a validation cache object using the
<code>pAllocator</code>. Applications <strong>can</strong> limit the
amount of data retrieved from a validation cache object in
<code>vkGetValidationCacheDataEXT</code>. Implementations
<strong>should</strong> not internally limit the total number of entries
added to a validation cache object or the total host memory
consumed.</p></td>
</tr>
</tbody>
</table>

Once created, a validation cache **can** be passed to the
`vkCreateShaderModule` command by adding this object to the
[VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html) structure’s
`pNext` chain. If a
[VkShaderModuleValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleValidationCacheCreateInfoEXT.html)
object is included in the
[VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html)::`pNext`
chain, and its `validationCache` field is not
[VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the implementation will query it
for possible reuse opportunities and update it with new content. The use
of the validation cache object in these commands is internally
synchronized, and the same validation cache object **can** be used in
multiple threads simultaneously.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Implementations <strong>should</strong> make every effort to limit
any critical sections to the actual accesses to the cache, which is
expected to be significantly shorter than the duration of the
<code>vkCreateShaderModule</code> command.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-vkCreateValidationCacheEXT-device-parameter"
  id="VUID-vkCreateValidationCacheEXT-device-parameter"></a>
  VUID-vkCreateValidationCacheEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateValidationCacheEXT-pCreateInfo-parameter"
  id="VUID-vkCreateValidationCacheEXT-pCreateInfo-parameter"></a>
  VUID-vkCreateValidationCacheEXT-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheCreateInfoEXT.html)
  structure

- <a href="#VUID-vkCreateValidationCacheEXT-pAllocator-parameter"
  id="VUID-vkCreateValidationCacheEXT-pAllocator-parameter"></a>
  VUID-vkCreateValidationCacheEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateValidationCacheEXT-pValidationCache-parameter"
  id="VUID-vkCreateValidationCacheEXT-pValidationCache-parameter"></a>
  VUID-vkCreateValidationCacheEXT-pValidationCache-parameter  
  `pValidationCache` **must** be a valid pointer to a
  [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheEXT.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_validation_cache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_validation_cache.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheCreateInfoEXT.html),
[VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateValidationCacheEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
