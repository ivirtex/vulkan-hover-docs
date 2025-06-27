# VkSystemAllocationScope(3) Manual Page

## Name

VkSystemAllocationScope - Allocation scope



## <a href="#_c_specification" class="anchor"></a>C Specification

Each allocation has an *allocation scope* defining its lifetime and
which object it is associated with. Possible values passed to the
`allocationScope` parameter of the callback functions specified by
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html), indicating the
allocation scope, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkSystemAllocationScope {
    VK_SYSTEM_ALLOCATION_SCOPE_COMMAND = 0,
    VK_SYSTEM_ALLOCATION_SCOPE_OBJECT = 1,
    VK_SYSTEM_ALLOCATION_SCOPE_CACHE = 2,
    VK_SYSTEM_ALLOCATION_SCOPE_DEVICE = 3,
    VK_SYSTEM_ALLOCATION_SCOPE_INSTANCE = 4,
} VkSystemAllocationScope;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_SYSTEM_ALLOCATION_SCOPE_COMMAND` specifies that the allocation is
  scoped to the duration of the Vulkan command.

- `VK_SYSTEM_ALLOCATION_SCOPE_OBJECT` specifies that the allocation is
  scoped to the lifetime of the Vulkan object that is being created or
  used.

- `VK_SYSTEM_ALLOCATION_SCOPE_CACHE` specifies that the allocation is
  scoped to the lifetime of a `VkPipelineCache` or
  `VkValidationCacheEXT` object.

- `VK_SYSTEM_ALLOCATION_SCOPE_DEVICE` specifies that the allocation is
  scoped to the lifetime of the Vulkan device.

- `VK_SYSTEM_ALLOCATION_SCOPE_INSTANCE` specifies that the allocation is
  scoped to the lifetime of the Vulkan instance.

Most Vulkan commands operate on a single object, or there is a sole
object that is being created or manipulated. When an allocation uses an
allocation scope of `VK_SYSTEM_ALLOCATION_SCOPE_OBJECT` or
`VK_SYSTEM_ALLOCATION_SCOPE_CACHE`, the allocation is scoped to the
object being created or manipulated.

When an implementation requires host memory, it will make callbacks to
the application using the most specific allocator and allocation scope
available:

- If an allocation is scoped to the duration of a command, the allocator
  will use the `VK_SYSTEM_ALLOCATION_SCOPE_COMMAND` allocation scope.
  The most specific allocator available is used: if the object being
  created or manipulated has an allocator, that object’s allocator will
  be used, else if the parent `VkDevice` has an allocator it will be
  used, else if the parent `VkInstance` has an allocator it will be
  used. Else,

- If an allocation is associated with a `VkValidationCacheEXT` or
  `VkPipelineCache` object, the allocator will use the
  `VK_SYSTEM_ALLOCATION_SCOPE_CACHE` allocation scope. The most specific
  allocator available is used (cache, else device, else instance). Else,

- If an allocation is scoped to the lifetime of an object, that object
  is being created or manipulated by the command, and that object’s type
  is not `VkDevice` or `VkInstance`, the allocator will use an
  allocation scope of `VK_SYSTEM_ALLOCATION_SCOPE_OBJECT`. The most
  specific allocator available is used (object, else device, else
  instance). Else,

- If an allocation is scoped to the lifetime of a device, the allocator
  will use an allocation scope of `VK_SYSTEM_ALLOCATION_SCOPE_DEVICE`.
  The most specific allocator available is used (device, else instance).
  Else,

- If the allocation is scoped to the lifetime of an instance and the
  instance has an allocator, its allocator will be used with an
  allocation scope of `VK_SYSTEM_ALLOCATION_SCOPE_INSTANCE`.

- Otherwise an implementation will allocate memory through an
  alternative mechanism that is unspecified.

## <a href="#_see_also" class="anchor"></a>See Also

[PFN_vkAllocationFunction](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkAllocationFunction.html),
[PFN_vkInternalAllocationNotification](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkInternalAllocationNotification.html),
[PFN_vkInternalFreeNotification](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkInternalFreeNotification.html),
[PFN_vkReallocationFunction](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkReallocationFunction.html),
[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSystemAllocationScope"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
