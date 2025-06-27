# vkSetEvent(3) Manual Page

## Name

vkSetEvent - Set an event to signaled state



## <a href="#_c_specification" class="anchor"></a>C Specification

To set the state of an event to signaled from the host, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkSetEvent(
    VkDevice                                    device,
    VkEvent                                     event);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the event.

- `event` is the event to set.

## <a href="#_description" class="anchor"></a>Description

When [vkSetEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetEvent.html) is executed on the host, it defines
an *event signal operation* which sets the event to the signaled state.

If `event` is already in the signaled state when
[vkSetEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetEvent.html) is executed, then
[vkSetEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetEvent.html) has no effect, and no event signal
operation occurs.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>If a command buffer is waiting for an event to be signaled from the
host, the application must signal the event before submitting the
command buffer, as described in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-submission-progress"
target="_blank" rel="noopener">queue forward progress</a>
section.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkSetEvent-event-03941"
  id="VUID-vkSetEvent-event-03941"></a> VUID-vkSetEvent-event-03941  
  `event` **must** not have been created with
  `VK_EVENT_CREATE_DEVICE_ONLY_BIT`

- <a href="#VUID-vkSetEvent-event-09543"
  id="VUID-vkSetEvent-event-09543"></a> VUID-vkSetEvent-event-09543  
  `event` **must** not be waited on by a command buffer in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">pending state</a>

Valid Usage (Implicit)

- <a href="#VUID-vkSetEvent-device-parameter"
  id="VUID-vkSetEvent-device-parameter"></a>
  VUID-vkSetEvent-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkSetEvent-event-parameter"
  id="VUID-vkSetEvent-event-parameter"></a>
  VUID-vkSetEvent-event-parameter  
  `event` **must** be a valid [VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html) handle

- <a href="#VUID-vkSetEvent-event-parent"
  id="VUID-vkSetEvent-event-parent"></a> VUID-vkSetEvent-event-parent  
  `event` **must** have been created, allocated, or retrieved from
  `device`

Host Synchronization

- Host access to `event` **must** be externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEvent.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkSetEvent"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
