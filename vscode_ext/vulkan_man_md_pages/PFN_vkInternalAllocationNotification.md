# PFN_vkInternalAllocationNotification(3) Manual Page

## Name

PFN_vkInternalAllocationNotification - Application-defined memory
allocation notification function



## <a href="#_c_specification" class="anchor"></a>C Specification

The type of `pfnInternalAllocation` is:

``` c
// Provided by VK_VERSION_1_0
typedef void (VKAPI_PTR *PFN_vkInternalAllocationNotification)(
    void*                                       pUserData,
    size_t                                      size,
    VkInternalAllocationType                    allocationType,
    VkSystemAllocationScope                     allocationScope);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `pUserData` is the value specified for
  [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)::`pUserData` in
  the allocator specified by the application.

- `size` is the requested size of an allocation.

- `allocationType` is a
  [VkInternalAllocationType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInternalAllocationType.html) value
  specifying the requested type of an allocation.

- `allocationScope` is a
  [VkSystemAllocationScope](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSystemAllocationScope.html) value
  specifying the allocation scope of the lifetime of the allocation, as
  described <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-host-allocation-scope"
  target="_blank" rel="noopener">here</a>.

## <a href="#_description" class="anchor"></a>Description

This is a purely informational callback.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#PFN_vkInternalAllocationNotification"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
