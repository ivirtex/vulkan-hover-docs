# PFN_vkAllocationFunction(3) Manual Page

## Name

PFN_vkAllocationFunction - Application-defined memory allocation
function



## <a href="#_c_specification" class="anchor"></a>C Specification

The type of `pfnAllocation` is:

``` c
// Provided by VK_VERSION_1_0
typedef void* (VKAPI_PTR *PFN_vkAllocationFunction)(
    void*                                       pUserData,
    size_t                                      size,
    size_t                                      alignment,
    VkSystemAllocationScope                     allocationScope);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `pUserData` is the value specified for
  [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)::`pUserData` in
  the allocator specified by the application.

- `size` is the size in bytes of the requested allocation.

- `alignment` is the requested alignment of the allocation in bytes and
  **must** be a power of two.

- `allocationScope` is a
  [VkSystemAllocationScope](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSystemAllocationScope.html) value
  specifying the allocation scope of the lifetime of the allocation, as
  described <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-host-allocation-scope"
  target="_blank" rel="noopener">here</a>.

## <a href="#_description" class="anchor"></a>Description

If `pfnAllocation` is unable to allocate the requested memory, it
**must** return `NULL`. If the allocation was successful, it **must**
return a valid pointer to memory allocation containing at least `size`
bytes, and with the pointer value being a multiple of `alignment`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Correct Vulkan operation <strong>cannot</strong> be assumed if the
application does not follow these rules.</p>
<p>For example, <code>pfnAllocation</code> (or
<code>pfnReallocation</code>) could cause termination of running Vulkan
instance(s) on a failed allocation for debugging purposes, either
directly or indirectly. In these circumstances, it
<strong>cannot</strong> be assumed that any part of any affected <a
href="VkInstance.html">VkInstance</a> objects are going to operate
correctly (even <a href="vkDestroyInstance.html">vkDestroyInstance</a>),
and the application <strong>must</strong> ensure it cleans up properly
via other means (e.g. process termination).</p></td>
</tr>
</tbody>
</table>

If `pfnAllocation` returns `NULL`, and if the implementation is unable
to continue correct processing of the current command without the
requested allocation, it **must** treat this as a runtime error, and
generate `VK_ERROR_OUT_OF_HOST_MEMORY` at the appropriate time for the
command in which the condition was detected, as described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fundamentals-errorcodes"
target="_blank" rel="noopener">Return Codes</a>.

If the implementation is able to continue correct processing of the
current command without the requested allocation, then it **may** do so,
and **must** not generate `VK_ERROR_OUT_OF_HOST_MEMORY` as a result of
this failed allocation.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkSystemAllocationScope](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSystemAllocationScope.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#PFN_vkAllocationFunction"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
