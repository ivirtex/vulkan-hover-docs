# VkEvent(3) Manual Page

## Name

VkEvent - Opaque handle to an event object



## <a href="#_c_specification" class="anchor"></a>C Specification

Events are a synchronization primitive that **can** be used to insert a
fine-grained dependency between commands submitted to the same queue, or
between the host and a queue. Events **must** not be used to insert a
dependency between commands submitted to different queues. Events have
two states - signaled and unsignaled. An application **can** signal or
unsignal an event either on the host or on the device. A device **can**
be made to wait for an event to become signaled before executing further
operations. No command exists to wait for an event to become signaled on
the host, but the current state of an event **can** be queried.

Events are represented by `VkEvent` handles:

``` c
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkEvent)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_DEFINE_NON_DISPATCHABLE_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html),
[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkExportMetalSharedEventInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalSharedEventInfoEXT.html),
[vkCmdResetEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdResetEvent.html),
[vkCmdResetEvent2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdResetEvent2.html),
[vkCmdResetEvent2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdResetEvent2KHR.html),
[vkCmdSetEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent.html),
[vkCmdSetEvent2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent2.html),
[vkCmdSetEvent2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetEvent2KHR.html),
[vkCmdWaitEvents](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents.html),
[vkCmdWaitEvents2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents2.html),
[vkCmdWaitEvents2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents2KHR.html),
[vkCreateEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateEvent.html),
[vkDestroyEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyEvent.html),
[vkGetEventStatus](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetEventStatus.html),
[vkResetEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkResetEvent.html), [vkSetEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetEvent.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkEvent"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
