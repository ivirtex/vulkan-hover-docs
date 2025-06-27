# VkDeviceGroupRenderPassBeginInfo(3) Manual Page

## Name

VkDeviceGroupRenderPassBeginInfo - Set the initial device mask and
render areas for a render pass instance



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of
[VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html) or
[VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html) includes a
`VkDeviceGroupRenderPassBeginInfo` structure, then that structure
includes a device mask and set of render areas for the render pass
instance.

The `VkDeviceGroupRenderPassBeginInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkDeviceGroupRenderPassBeginInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           deviceMask;
    uint32_t           deviceRenderAreaCount;
    const VkRect2D*    pDeviceRenderAreas;
} VkDeviceGroupRenderPassBeginInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_device_group
typedef VkDeviceGroupRenderPassBeginInfo VkDeviceGroupRenderPassBeginInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `deviceMask` is the device mask for the render pass instance.

- `deviceRenderAreaCount` is the number of elements in the
  `pDeviceRenderAreas` array.

- `pDeviceRenderAreas` is a pointer to an array of
  [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html) structures defining the render area for each
  physical device.

## <a href="#_description" class="anchor"></a>Description

The `deviceMask` serves several purposes. It is an upper bound on the
set of physical devices that **can** be used during the render pass
instance, and the initial device mask when the render pass instance
begins. In addition, commands transitioning to the next subpass in a
render pass instance and commands ending the render pass instance, and,
accordingly render pass <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-load-operations"
target="_blank" rel="noopener">load</a>, <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-store-operations"
target="_blank" rel="noopener">store</a>, and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-resolve-operations"
target="_blank" rel="noopener">multisample resolve</a> operations and
subpass dependencies corresponding to the render pass instance, are
executed on the physical devices included in the device mask provided
here.

If `deviceRenderAreaCount` is not zero, then the elements of
`pDeviceRenderAreas` override the value of
[VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html)::`renderArea`, and
provide a render area specific to each physical device. These render
areas serve the same purpose as
[VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html)::`renderArea`,
including controlling the region of attachments that are cleared by
`VK_ATTACHMENT_LOAD_OP_CLEAR` and that are resolved into resolve
attachments.

If this structure is not present, the render pass instance’s device mask
is the value of
[VkDeviceGroupCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupCommandBufferBeginInfo.html)::`deviceMask`.
If this structure is not present or if `deviceRenderAreaCount` is zero,
[VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassBeginInfo.html)::`renderArea` is
used for all physical devices.

Valid Usage

- <a href="#VUID-VkDeviceGroupRenderPassBeginInfo-deviceMask-00905"
  id="VUID-VkDeviceGroupRenderPassBeginInfo-deviceMask-00905"></a>
  VUID-VkDeviceGroupRenderPassBeginInfo-deviceMask-00905  
  `deviceMask` **must** be a valid device mask value

- <a href="#VUID-VkDeviceGroupRenderPassBeginInfo-deviceMask-00906"
  id="VUID-VkDeviceGroupRenderPassBeginInfo-deviceMask-00906"></a>
  VUID-VkDeviceGroupRenderPassBeginInfo-deviceMask-00906  
  `deviceMask` **must** not be zero

- <a href="#VUID-VkDeviceGroupRenderPassBeginInfo-deviceMask-00907"
  id="VUID-VkDeviceGroupRenderPassBeginInfo-deviceMask-00907"></a>
  VUID-VkDeviceGroupRenderPassBeginInfo-deviceMask-00907  
  `deviceMask` **must** be a subset of the command buffer’s initial
  device mask

- <a
  href="#VUID-VkDeviceGroupRenderPassBeginInfo-deviceRenderAreaCount-00908"
  id="VUID-VkDeviceGroupRenderPassBeginInfo-deviceRenderAreaCount-00908"></a>
  VUID-VkDeviceGroupRenderPassBeginInfo-deviceRenderAreaCount-00908  
  `deviceRenderAreaCount` **must** either be zero or equal to the number
  of physical devices in the logical device

- <a href="#VUID-VkDeviceGroupRenderPassBeginInfo-offset-06166"
  id="VUID-VkDeviceGroupRenderPassBeginInfo-offset-06166"></a>
  VUID-VkDeviceGroupRenderPassBeginInfo-offset-06166  
  The `offset.x` member of any element of `pDeviceRenderAreas` **must**
  be greater than or equal to 0

- <a href="#VUID-VkDeviceGroupRenderPassBeginInfo-offset-06167"
  id="VUID-VkDeviceGroupRenderPassBeginInfo-offset-06167"></a>
  VUID-VkDeviceGroupRenderPassBeginInfo-offset-06167  
  The `offset.y` member of any element of `pDeviceRenderAreas` **must**
  be greater than or equal to 0

- <a href="#VUID-VkDeviceGroupRenderPassBeginInfo-offset-06168"
  id="VUID-VkDeviceGroupRenderPassBeginInfo-offset-06168"></a>
  VUID-VkDeviceGroupRenderPassBeginInfo-offset-06168  
  The sum of the `offset.x` and `extent.width` members of any element of
  `pDeviceRenderAreas` **must** be less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxFramebufferWidth"
  target="_blank" rel="noopener"><code>maxFramebufferWidth</code></a>

- <a href="#VUID-VkDeviceGroupRenderPassBeginInfo-offset-06169"
  id="VUID-VkDeviceGroupRenderPassBeginInfo-offset-06169"></a>
  VUID-VkDeviceGroupRenderPassBeginInfo-offset-06169  
  The sum of the `offset.y` and `extent.height` members of any element
  of `pDeviceRenderAreas` **must** be less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxFramebufferHeight"
  target="_blank" rel="noopener"><code>maxFramebufferHeight</code></a>

- <a href="#VUID-VkDeviceGroupRenderPassBeginInfo-extent-08998"
  id="VUID-VkDeviceGroupRenderPassBeginInfo-extent-08998"></a>
  VUID-VkDeviceGroupRenderPassBeginInfo-extent-08998  
  The `extent.width` member of any element of `pDeviceRenderAreas`
  **must** be greater than 0

- <a href="#VUID-VkDeviceGroupRenderPassBeginInfo-extent-08999"
  id="VUID-VkDeviceGroupRenderPassBeginInfo-extent-08999"></a>
  VUID-VkDeviceGroupRenderPassBeginInfo-extent-08999  
  The `extent.height` member of any element of `pDeviceRenderAreas`
  **must** be greater than 0

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceGroupRenderPassBeginInfo-sType-sType"
  id="VUID-VkDeviceGroupRenderPassBeginInfo-sType-sType"></a>
  VUID-VkDeviceGroupRenderPassBeginInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEVICE_GROUP_RENDER_PASS_BEGIN_INFO`

- <a
  href="#VUID-VkDeviceGroupRenderPassBeginInfo-pDeviceRenderAreas-parameter"
  id="VUID-VkDeviceGroupRenderPassBeginInfo-pDeviceRenderAreas-parameter"></a>
  VUID-VkDeviceGroupRenderPassBeginInfo-pDeviceRenderAreas-parameter  
  If `deviceRenderAreaCount` is not `0`, `pDeviceRenderAreas` **must**
  be a valid pointer to an array of `deviceRenderAreaCount`
  [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html) structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceGroupRenderPassBeginInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
