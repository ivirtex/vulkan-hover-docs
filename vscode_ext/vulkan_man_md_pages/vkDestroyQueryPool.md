# vkDestroyQueryPool(3) Manual Page

## Name

vkDestroyQueryPool - Destroy a query pool object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a query pool, call:

``` c
// Provided by VK_VERSION_1_0
void vkDestroyQueryPool(
    VkDevice                                    device,
    VkQueryPool                                 queryPool,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the query pool.

- `queryPool` is the query pool to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroyQueryPool-queryPool-00793"
  id="VUID-vkDestroyQueryPool-queryPool-00793"></a>
  VUID-vkDestroyQueryPool-queryPool-00793  
  All submitted commands that refer to `queryPool` **must** have
  completed execution

- <a href="#VUID-vkDestroyQueryPool-queryPool-00794"
  id="VUID-vkDestroyQueryPool-queryPool-00794"></a>
  VUID-vkDestroyQueryPool-queryPool-00794  
  If `VkAllocationCallbacks` were provided when `queryPool` was created,
  a compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyQueryPool-queryPool-00795"
  id="VUID-vkDestroyQueryPool-queryPool-00795"></a>
  VUID-vkDestroyQueryPool-queryPool-00795  
  If no `VkAllocationCallbacks` were provided when `queryPool` was
  created, `pAllocator` **must** be `NULL`

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Applications <strong>can</strong> verify that <code>queryPool</code>
<strong>can</strong> be destroyed by checking that
<code>vkGetQueryPoolResults</code>() without the
<code>VK_QUERY_RESULT_PARTIAL_BIT</code> flag returns
<code>VK_SUCCESS</code> for all queries that are used in command buffers
submitted for execution.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyQueryPool-device-parameter"
  id="VUID-vkDestroyQueryPool-device-parameter"></a>
  VUID-vkDestroyQueryPool-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyQueryPool-queryPool-parameter"
  id="VUID-vkDestroyQueryPool-queryPool-parameter"></a>
  VUID-vkDestroyQueryPool-queryPool-parameter  
  If `queryPool` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `queryPool` **must** be a valid [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html) handle

- <a href="#VUID-vkDestroyQueryPool-pAllocator-parameter"
  id="VUID-vkDestroyQueryPool-pAllocator-parameter"></a>
  VUID-vkDestroyQueryPool-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyQueryPool-queryPool-parent"
  id="VUID-vkDestroyQueryPool-queryPool-parent"></a>
  VUID-vkDestroyQueryPool-queryPool-parent  
  If `queryPool` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `queryPool` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPool.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyQueryPool"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
