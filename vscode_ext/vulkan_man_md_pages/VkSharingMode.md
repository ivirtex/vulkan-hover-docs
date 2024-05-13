# VkSharingMode(3) Manual Page

## Name

VkSharingMode - Buffer and image sharing modes



## <a href="#_c_specification" class="anchor"></a>C Specification

Buffer and image objects are created with a *sharing mode* controlling
how they **can** be accessed from queues. The supported sharing modes
are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkSharingMode {
    VK_SHARING_MODE_EXCLUSIVE = 0,
    VK_SHARING_MODE_CONCURRENT = 1,
} VkSharingMode;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_SHARING_MODE_EXCLUSIVE` specifies that access to any range or
  image subresource of the object will be exclusive to a single queue
  family at a time.

- `VK_SHARING_MODE_CONCURRENT` specifies that concurrent access to any
  range or image subresource of the object from multiple queue families
  is supported.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p><code>VK_SHARING_MODE_CONCURRENT</code> <strong>may</strong> result
in lower performance access to the buffer or image than
<code>VK_SHARING_MODE_EXCLUSIVE</code>.</p></td>
</tr>
</tbody>
</table>

Ranges of buffers and image subresources of image objects created using
`VK_SHARING_MODE_EXCLUSIVE` **must** only be accessed by queues in the
queue family that has *ownership* of the resource. Upon creation, such
resources are not owned by any queue family; ownership is implicitly
acquired upon first use within a queue. Once a resource using
`VK_SHARING_MODE_EXCLUSIVE` is owned by some queue family, the
application **must** perform a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-queue-transfers"
target="_blank" rel="noopener">queue family ownership transfer</a> to
make the memory contents of a range or image subresource accessible to a
different queue family.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Images still require a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-layouts"
target="_blank" rel="noopener">layout transition</a> from
<code>VK_IMAGE_LAYOUT_UNDEFINED</code> or
<code>VK_IMAGE_LAYOUT_PREINITIALIZED</code> before being used on the
first queue.</p></td>
</tr>
</tbody>
</table>

A queue family **can** take ownership of an image subresource or buffer
range of a resource created with `VK_SHARING_MODE_EXCLUSIVE`, without an
ownership transfer, in the same way as for a resource that was just
created; however, taking ownership in this way has the effect that the
contents of the image subresource or buffer range are undefined.

Ranges of buffers and image subresources of image objects created using
`VK_SHARING_MODE_CONCURRENT` **must** only be accessed by queues from
the queue families specified through the `queueFamilyIndexCount` and
`pQueueFamilyIndices` members of the corresponding create info
structures.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html),
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html),
[VkPhysicalDeviceImageDrmFormatModifierInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageDrmFormatModifierInfoEXT.html),
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSharingMode"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
