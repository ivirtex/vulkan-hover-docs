# PFN_vkFreeFunction(3) Manual Page

## Name

PFN_vkFreeFunction - Application-defined memory free function



## <a href="#_c_specification" class="anchor"></a>C Specification

The type of `pfnFree` is:

``` c
// Provided by VK_VERSION_1_0
typedef void (VKAPI_PTR *PFN_vkFreeFunction)(
    void*                                       pUserData,
    void*                                       pMemory);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `pUserData` is the value specified for
  [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)::`pUserData` in
  the allocator specified by the application.

- `pMemory` is the allocation to be freed.

## <a href="#_description" class="anchor"></a>Description

`pMemory` **may** be `NULL`, which the callback **must** handle safely.
If `pMemory` is non-`NULL`, it **must** be a pointer previously
allocated by `pfnAllocation` or `pfnReallocation`. The application
**should** free this memory.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#PFN_vkFreeFunction"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
