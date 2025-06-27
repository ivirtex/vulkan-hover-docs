# VkRenderPassBeginInfo(3) Manual Page

## Name

VkRenderPassBeginInfo - Structure specifying render pass begin
information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkRenderPassBeginInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkRenderPassBeginInfo {
    VkStructureType        sType;
    const void*            pNext;
    VkRenderPass           renderPass;
    VkFramebuffer          framebuffer;
    VkRect2D               renderArea;
    uint32_t               clearValueCount;
    const VkClearValue*    pClearValues;
} VkRenderPassBeginInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `renderPass` is the render pass to begin an instance of.

- `framebuffer` is the framebuffer containing the attachments that are
  used with the render pass.

- `renderArea` is the render area that is affected by the render pass
  instance, and is described in more detail below.

- `clearValueCount` is the number of elements in `pClearValues`.

- `pClearValues` is a pointer to an array of `clearValueCount`
  [VkClearValue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearValue.html) structures containing clear values
  for each attachment, if the attachment uses a `loadOp` value of
  `VK_ATTACHMENT_LOAD_OP_CLEAR` or if the attachment has a depth/stencil
  format and uses a `stencilLoadOp` value of
  `VK_ATTACHMENT_LOAD_OP_CLEAR`. The array is indexed by attachment
  number. Only elements corresponding to cleared attachments are used.
  Other elements of `pClearValues` are ignored.

## <a href="#_description" class="anchor"></a>Description

`renderArea` is the render area that is affected by the render pass
instance. The effects of attachment load, store and multisample resolve
operations are restricted to the pixels whose x and y coordinates fall
within the render area on all attachments. The render area extends to
all layers of `framebuffer`. The application **must** ensure (using
scissor if necessary) that all rendering is contained within the render
area. The render area, after any transform specified by
[VkRenderPassTransformBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassTransformBeginInfoQCOM.html)::`transform`
is applied, **must** be contained within the framebuffer dimensions.

If <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vertexpostproc-renderpass-transform"
target="_blank" rel="noopener">render pass transform</a> is enabled,
then `renderArea` **must** equal the framebuffer pre-transformed
dimensions. After `renderArea` has been transformed by
[VkRenderPassTransformBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassTransformBeginInfoQCOM.html)::`transform`,
the resulting render area **must** be equal to the framebuffer
dimensions.

If multiview is enabled in `renderPass`, and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview-per-view-render-areas"
target="_blank"
rel="noopener"><code>multiviewPerViewRenderAreas</code></a> feature is
enabled, and there is an instance of
[VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM.html)
included in the `pNext` chain with `perViewRenderAreaCount` not equal to
`0`, then the elements of
[VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM.html)::`pPerViewRenderAreas`
override `renderArea` and define a render area for each view. In this
case, `renderArea` **must** be set to an area at least as large as the
union of all the per-view render areas.

If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-subpassShading"
target="_blank" rel="noopener"><code>subpassShading</code></a> feature
is enabled, then `renderArea` **must** equal the framebuffer dimensions.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>There <strong>may</strong> be a performance cost for using a render
area smaller than the framebuffer, unless it matches the render area
granularity for the render pass.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkRenderPassBeginInfo-clearValueCount-00902"
  id="VUID-VkRenderPassBeginInfo-clearValueCount-00902"></a>
  VUID-VkRenderPassBeginInfo-clearValueCount-00902  
  `clearValueCount` **must** be greater than the largest attachment
  index in `renderPass` specifying a `loadOp` (or `stencilLoadOp`, if
  the attachment has a depth/stencil format) of
  `VK_ATTACHMENT_LOAD_OP_CLEAR`

- <a href="#VUID-VkRenderPassBeginInfo-clearValueCount-04962"
  id="VUID-VkRenderPassBeginInfo-clearValueCount-04962"></a>
  VUID-VkRenderPassBeginInfo-clearValueCount-04962  
  If `clearValueCount` is not `0`, `pClearValues` **must** be a valid
  pointer to an array of `clearValueCount`
  [VkClearValue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearValue.html) unions

- <a href="#VUID-VkRenderPassBeginInfo-renderPass-00904"
  id="VUID-VkRenderPassBeginInfo-renderPass-00904"></a>
  VUID-VkRenderPassBeginInfo-renderPass-00904  
  `renderPass` **must** be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-compatibility"
  target="_blank" rel="noopener">compatible</a> with the `renderPass`
  member of the [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html)
  structure specified when creating `framebuffer`

- <a href="#VUID-VkRenderPassBeginInfo-None-08996"
  id="VUID-VkRenderPassBeginInfo-None-08996"></a>
  VUID-VkRenderPassBeginInfo-None-08996  
  If the `pNext` chain does not contain
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)
  or its `deviceRenderAreaCount` member is equal to 0,
  `renderArea.extent.width` **must** be greater than 0

- <a href="#VUID-VkRenderPassBeginInfo-None-08997"
  id="VUID-VkRenderPassBeginInfo-None-08997"></a>
  VUID-VkRenderPassBeginInfo-None-08997  
  If the `pNext` chain does not contain
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)
  or its `deviceRenderAreaCount` member is equal to 0,
  `renderArea.extent.height` **must** be greater than 0

- <a href="#VUID-VkRenderPassBeginInfo-pNext-02850"
  id="VUID-VkRenderPassBeginInfo-pNext-02850"></a>
  VUID-VkRenderPassBeginInfo-pNext-02850  
  If the `pNext` chain does not contain
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)
  or its `deviceRenderAreaCount` member is equal to 0,
  `renderArea.offset.x` **must** be greater than or equal to 0

- <a href="#VUID-VkRenderPassBeginInfo-pNext-02851"
  id="VUID-VkRenderPassBeginInfo-pNext-02851"></a>
  VUID-VkRenderPassBeginInfo-pNext-02851  
  If the `pNext` chain does not contain
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)
  or its `deviceRenderAreaCount` member is equal to 0,
  `renderArea.offset.y` **must** be greater than or equal to 0

- <a href="#VUID-VkRenderPassBeginInfo-pNext-02852"
  id="VUID-VkRenderPassBeginInfo-pNext-02852"></a>
  VUID-VkRenderPassBeginInfo-pNext-02852  
  If the `pNext` chain does not contain
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)
  or its `deviceRenderAreaCount` member is equal to 0,
  `renderArea.offset.x` + `renderArea.extent.width` **must** be less
  than or equal to
  [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html)::`width` the
  `framebuffer` was created with

- <a href="#VUID-VkRenderPassBeginInfo-pNext-02853"
  id="VUID-VkRenderPassBeginInfo-pNext-02853"></a>
  VUID-VkRenderPassBeginInfo-pNext-02853  
  If the `pNext` chain does not contain
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)
  or its `deviceRenderAreaCount` member is equal to 0,
  `renderArea.offset.y` + `renderArea.extent.height` **must** be less
  than or equal to
  [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html)::`height` the
  `framebuffer` was created with

- <a href="#VUID-VkRenderPassBeginInfo-pNext-02856"
  id="VUID-VkRenderPassBeginInfo-pNext-02856"></a>
  VUID-VkRenderPassBeginInfo-pNext-02856  
  If the `pNext` chain contains
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html),
  `offset.x` + `extent.width` of each element of `pDeviceRenderAreas`
  **must** be less than or equal to
  [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html)::`width` the
  `framebuffer` was created with

- <a href="#VUID-VkRenderPassBeginInfo-pNext-02857"
  id="VUID-VkRenderPassBeginInfo-pNext-02857"></a>
  VUID-VkRenderPassBeginInfo-pNext-02857  
  If the `pNext` chain contains
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html),
  `offset.y` + `extent.height` of each element of `pDeviceRenderAreas`
  **must** be less than or equal to
  [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html)::`height` the
  `framebuffer` was created with

- <a href="#VUID-VkRenderPassBeginInfo-framebuffer-03207"
  id="VUID-VkRenderPassBeginInfo-framebuffer-03207"></a>
  VUID-VkRenderPassBeginInfo-framebuffer-03207  
  If `framebuffer` was created with a
  [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html)::`flags` value
  that did not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, and the
  `pNext` chain includes a
  [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassAttachmentBeginInfo.html)
  structure, its `attachmentCount` **must** be zero

- <a href="#VUID-VkRenderPassBeginInfo-framebuffer-03208"
  id="VUID-VkRenderPassBeginInfo-framebuffer-03208"></a>
  VUID-VkRenderPassBeginInfo-framebuffer-03208  
  If `framebuffer` was created with a
  [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html)::`flags` value
  that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the
  `attachmentCount` of a
  [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassAttachmentBeginInfo.html)
  structure included in the `pNext` chain **must** be equal to the value
  of
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)::`attachmentImageInfoCount`
  used to create `framebuffer`

- <a href="#VUID-VkRenderPassBeginInfo-framebuffer-02780"
  id="VUID-VkRenderPassBeginInfo-framebuffer-02780"></a>
  VUID-VkRenderPassBeginInfo-framebuffer-02780  
  If `framebuffer` was created with a
  [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html)::`flags` value
  that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of
  the `pAttachments` member of a
  [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassAttachmentBeginInfo.html)
  structure included in the `pNext` chain **must** have been created on
  the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) as `framebuffer` and `renderPass`

- <a href="#VUID-VkRenderPassBeginInfo-framebuffer-03209"
  id="VUID-VkRenderPassBeginInfo-framebuffer-03209"></a>
  VUID-VkRenderPassBeginInfo-framebuffer-03209  
  If `framebuffer` was created with a
  [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html)::`flags` value
  that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of
  the `pAttachments` member of a
  [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassAttachmentBeginInfo.html)
  structure included in the `pNext` chain **must** be a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) of an image created with a value of
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`flags` equal to the
  `flags` member of the corresponding element of
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)::`pAttachmentImageInfos`
  used to create `framebuffer`

- <a href="#VUID-VkRenderPassBeginInfo-framebuffer-04627"
  id="VUID-VkRenderPassBeginInfo-framebuffer-04627"></a>
  VUID-VkRenderPassBeginInfo-framebuffer-04627  
  If `framebuffer` was created with a
  [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html)::`flags` value
  that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of
  the `pAttachments` member of a
  [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassAttachmentBeginInfo.html)
  structure included in the `pNext` chain **must** be a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-inherited-usage"
  target="_blank" rel="noopener">an inherited usage</a> equal to the
  `usage` member of the corresponding element of
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)::`pAttachmentImageInfos`
  used to create `framebuffer`

- <a href="#VUID-VkRenderPassBeginInfo-framebuffer-03211"
  id="VUID-VkRenderPassBeginInfo-framebuffer-03211"></a>
  VUID-VkRenderPassBeginInfo-framebuffer-03211  
  If `framebuffer` was created with a
  [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html)::`flags` value
  that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of
  the `pAttachments` member of a
  [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassAttachmentBeginInfo.html)
  structure included in the `pNext` chain **must** be a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) with a width equal to the `width`
  member of the corresponding element of
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)::`pAttachmentImageInfos`
  used to create `framebuffer`

- <a href="#VUID-VkRenderPassBeginInfo-framebuffer-03212"
  id="VUID-VkRenderPassBeginInfo-framebuffer-03212"></a>
  VUID-VkRenderPassBeginInfo-framebuffer-03212  
  If `framebuffer` was created with a
  [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html)::`flags` value
  that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of
  the `pAttachments` member of a
  [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassAttachmentBeginInfo.html)
  structure included in the `pNext` chain **must** be a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) with a height equal to the `height`
  member of the corresponding element of
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)::`pAttachmentImageInfos`
  used to create `framebuffer`

- <a href="#VUID-VkRenderPassBeginInfo-framebuffer-03213"
  id="VUID-VkRenderPassBeginInfo-framebuffer-03213"></a>
  VUID-VkRenderPassBeginInfo-framebuffer-03213  
  If `framebuffer` was created with a
  [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html)::`flags` value
  that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of
  the `pAttachments` member of a
  [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassAttachmentBeginInfo.html)
  structure included in the `pNext` chain **must** be a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) of an image created with a value of
  [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html)::`subresourceRange.layerCount`
  equal to the `layerCount` member of the corresponding element of
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)::`pAttachmentImageInfos`
  used to create `framebuffer`

- <a href="#VUID-VkRenderPassBeginInfo-framebuffer-03214"
  id="VUID-VkRenderPassBeginInfo-framebuffer-03214"></a>
  VUID-VkRenderPassBeginInfo-framebuffer-03214  
  If `framebuffer` was created with a
  [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html)::`flags` value
  that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of
  the `pAttachments` member of a
  [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassAttachmentBeginInfo.html)
  structure included in the `pNext` chain **must** be a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) of an image created with a value of
  [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html)::`viewFormatCount`
  equal to the `viewFormatCount` member of the corresponding element of
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)::`pAttachmentImageInfos`
  used to create `framebuffer`

- <a href="#VUID-VkRenderPassBeginInfo-framebuffer-03215"
  id="VUID-VkRenderPassBeginInfo-framebuffer-03215"></a>
  VUID-VkRenderPassBeginInfo-framebuffer-03215  
  If `framebuffer` was created with a
  [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html)::`flags` value
  that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of
  the `pAttachments` member of a
  [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassAttachmentBeginInfo.html)
  structure included in the `pNext` chain **must** be a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) of an image created with a set of
  elements in
  [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html)::`pViewFormats`
  equal to the set of elements in the `pViewFormats` member of the
  corresponding element of
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)::`pAttachmentImageInfos`
  used to create `framebuffer`

- <a href="#VUID-VkRenderPassBeginInfo-framebuffer-03216"
  id="VUID-VkRenderPassBeginInfo-framebuffer-03216"></a>
  VUID-VkRenderPassBeginInfo-framebuffer-03216  
  If `framebuffer` was created with a
  [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html)::`flags` value
  that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of
  the `pAttachments` member of a
  [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassAttachmentBeginInfo.html)
  structure included in the `pNext` chain **must** be a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) of an image created with a value of
  [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html)::`format` equal to
  the corresponding value of
  [VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription.html)::`format` in
  `renderPass`

- <a href="#VUID-VkRenderPassBeginInfo-framebuffer-09353"
  id="VUID-VkRenderPassBeginInfo-framebuffer-09353"></a>
  VUID-VkRenderPassBeginInfo-framebuffer-09353  
  If `framebuffer` was created with a
  [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html)::`flags` value
  that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve"
  target="_blank"
  rel="noopener"><code>nullColorAttachmentWithExternalFormatResolve</code></a>
  is `VK_FALSE`, the format of the color attachment for each subpass
  that includes an external format image as a resolve attachment
  **must** have a format equal to the value of
  [VkAndroidHardwareBufferFormatResolvePropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferFormatResolvePropertiesANDROID.html)::`colorAttachmentFormat`
  as returned by a call to
  [vkGetAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAndroidHardwareBufferPropertiesANDROID.html)
  for the Android hardware buffer that was used to create the image view
  use as its resolve attachment

- <a href="#VUID-VkRenderPassBeginInfo-framebuffer-09354"
  id="VUID-VkRenderPassBeginInfo-framebuffer-09354"></a>
  VUID-VkRenderPassBeginInfo-framebuffer-09354  
  If `framebuffer` was created with a
  [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html)::`flags` value
  that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of
  the `pAttachments` member of a
  [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassAttachmentBeginInfo.html)
  structure included in the `pNext` chain **must** be a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) of an image created with a value of
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  equal to
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  in the `pNext` chain of the corresponding
  [VkAttachmentDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription2.html) structure
  used to create `renderPass`

- <a href="#VUID-VkRenderPassBeginInfo-framebuffer-09047"
  id="VUID-VkRenderPassBeginInfo-framebuffer-09047"></a>
  VUID-VkRenderPassBeginInfo-framebuffer-09047  
  If `framebuffer` was created with a
  [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html)::`flags` value
  that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of
  the `pAttachments` member of a
  [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassAttachmentBeginInfo.html)
  structure included in the `pNext` chain **must** be a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) of an image created with a value of
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`samples` equal to the
  corresponding value of
  [VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription.html)::`samples` in
  `renderPass` , or `VK_SAMPLE_COUNT_1_BIT` if `renderPass` was created
  with
  [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html)
  structure in the `pNext` chain with
  `multisampledRenderToSingleSampledEnable` equal to `VK_TRUE`

- <a href="#VUID-VkRenderPassBeginInfo-pNext-02869"
  id="VUID-VkRenderPassBeginInfo-pNext-02869"></a>
  VUID-VkRenderPassBeginInfo-pNext-02869  
  If the `pNext` chain includes
  [VkRenderPassTransformBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassTransformBeginInfoQCOM.html),
  `renderArea.offset` **must** equal (0,0)

- <a href="#VUID-VkRenderPassBeginInfo-pNext-02870"
  id="VUID-VkRenderPassBeginInfo-pNext-02870"></a>
  VUID-VkRenderPassBeginInfo-pNext-02870  
  If the `pNext` chain includes
  [VkRenderPassTransformBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassTransformBeginInfoQCOM.html),
  `renderArea.extent` transformed by
  [VkRenderPassTransformBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassTransformBeginInfoQCOM.html)::`transform`
  **must** equal the `framebuffer` dimensions

- <a href="#VUID-VkRenderPassBeginInfo-perViewRenderAreaCount-07859"
  id="VUID-VkRenderPassBeginInfo-perViewRenderAreaCount-07859"></a>
  VUID-VkRenderPassBeginInfo-perViewRenderAreaCount-07859  
  If the `perViewRenderAreaCount` member of a
  [VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM.html)
  structure included in the `pNext` chain is not `0`, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview-per-view-render-areas"
  target="_blank"
  rel="noopener"><code>multiviewPerViewRenderAreas</code></a> feature
  **must** be enabled

- <a href="#VUID-VkRenderPassBeginInfo-perViewRenderAreaCount-07860"
  id="VUID-VkRenderPassBeginInfo-perViewRenderAreaCount-07860"></a>
  VUID-VkRenderPassBeginInfo-perViewRenderAreaCount-07860  
  If the `perViewRenderAreaCount` member of a
  [VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM.html)
  structure included in the `pNext` chain is not `0`, then `renderArea`
  **must** specify a render area that includes the union of all per view
  render areas

- <a href="#VUID-VkRenderPassBeginInfo-pNext-09539"
  id="VUID-VkRenderPassBeginInfo-pNext-09539"></a>
  VUID-VkRenderPassBeginInfo-pNext-09539  
  If the `pNext` chain contains a
  [VkRenderPassStripeBeginInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassStripeBeginInfoARM.html)
  structure, the union of stripe areas defined by the elements of
  [VkRenderPassStripeInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassStripeInfoARM.html)::`pStripeInfos`
  **must** cover the `renderArea`

Valid Usage (Implicit)

- <a href="#VUID-VkRenderPassBeginInfo-sType-sType"
  id="VUID-VkRenderPassBeginInfo-sType-sType"></a>
  VUID-VkRenderPassBeginInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO`

- <a href="#VUID-VkRenderPassBeginInfo-pNext-pNext"
  id="VUID-VkRenderPassBeginInfo-pNext-pNext"></a>
  VUID-VkRenderPassBeginInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html),
  [VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM.html),
  [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassAttachmentBeginInfo.html),
  [VkRenderPassSampleLocationsBeginInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassSampleLocationsBeginInfoEXT.html),
  [VkRenderPassStripeBeginInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassStripeBeginInfoARM.html),
  or
  [VkRenderPassTransformBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassTransformBeginInfoQCOM.html)

- <a href="#VUID-VkRenderPassBeginInfo-sType-unique"
  id="VUID-VkRenderPassBeginInfo-sType-unique"></a>
  VUID-VkRenderPassBeginInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkRenderPassBeginInfo-renderPass-parameter"
  id="VUID-VkRenderPassBeginInfo-renderPass-parameter"></a>
  VUID-VkRenderPassBeginInfo-renderPass-parameter  
  `renderPass` **must** be a valid [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html)
  handle

- <a href="#VUID-VkRenderPassBeginInfo-framebuffer-parameter"
  id="VUID-VkRenderPassBeginInfo-framebuffer-parameter"></a>
  VUID-VkRenderPassBeginInfo-framebuffer-parameter  
  `framebuffer` **must** be a valid [VkFramebuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebuffer.html)
  handle

- <a href="#VUID-VkRenderPassBeginInfo-commonparent"
  id="VUID-VkRenderPassBeginInfo-commonparent"></a>
  VUID-VkRenderPassBeginInfo-commonparent  
  Both of `framebuffer`, and `renderPass` **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkClearValue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkClearValue.html), [VkFramebuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebuffer.html),
[VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html), [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdBeginRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass.html),
[vkCmdBeginRenderPass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass2.html),
[vkCmdBeginRenderPass2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderPass2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderPassBeginInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
