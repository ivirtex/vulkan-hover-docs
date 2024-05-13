# VK_EXT_image_drm_format_modifier(3) Manual Page

## Name

VK_EXT_image_drm_format_modifier - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

159

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

             [VK_KHR_bind_memory2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_bind_memory2.html)  
             and  
            
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
             and  
            
[VK_KHR_sampler_ycbcr_conversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_sampler_ycbcr_conversion.html)  
         or  
         [Version 1.1](#versions-1.1)  
     and  
     [VK_KHR_image_format_list](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_image_format_list.html)  
or  
[Version 1.2](#versions-1.2)  

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_VERSION_1_3

- Interacts with VK_KHR_format_feature_flags2

## <a href="#_contact" class="anchor"></a>Contact

- Lina Versace <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_image_drm_format_modifier%5D%20@versalinyaa%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_image_drm_format_modifier%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>versalinyaa</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-09-30

**IP Status**  
No known IP claims.

**Contributors**  
- Antoine Labour, Google

- Bas Nieuwenhuizen, Google

- Lina Versace, Google

- James Jones, NVIDIA

- Faith Ekstrand, Intel

- Jőrg Wagner, ARM

- Kristian Høgsberg Kristensen, Google

- Ray Smith, ARM

## <a href="#_description" class="anchor"></a>Description

This extension provides the ability to use *DRM format modifiers* with
images, enabling Vulkan to better integrate with the Linux ecosystem of
graphics, video, and display APIs.

Its functionality closely overlaps with
`EGL_EXT_image_dma_buf_import_modifiers`

[2](https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_image_drm_format_modifier-fn2)^
and `EGL_MESA_image_dma_buf_export`

[3](https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_image_drm_format_modifier-fn3)^.
Unlike the EGL extensions, this extension does not require the use of a
specific handle type (such as a dma_buf) for external memory and
provides more explicit control of image creation.

## <a href="#_introduction_to_drm_format_modifiers" class="anchor"></a>Introduction to DRM Format Modifiers

A *DRM format modifier* is a 64-bit, vendor-prefixed, semi-opaque
unsigned integer. Most *modifiers* represent a concrete, vendor-specific
tiling format for images. Some exceptions are `DRM_FORMAT_MOD_LINEAR`
(which is not vendor-specific); `DRM_FORMAT_MOD_NONE` (which is an alias
of `DRM_FORMAT_MOD_LINEAR` due to historical accident); and
`DRM_FORMAT_MOD_INVALID` (which does not represent a tiling format). The
*modifier’s* vendor prefix consists of the 8 most significant bits. The
canonical list of *modifiers* and vendor prefixes is found in
[`drm_fourcc.h`](https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/include/uapi/drm/drm_fourcc.h)
in the Linux kernel source. The other dominant source of *modifiers* are
vendor kernel trees.

One goal of *modifiers* in the Linux ecosystem is to enumerate for each
vendor a reasonably sized set of tiling formats that are appropriate for
images shared across processes, APIs, and/or devices, where each
participating component may possibly be from different vendors. A
non-goal is to enumerate all tiling formats supported by all vendors.
Some tiling formats used internally by vendors are inappropriate for
sharing; no *modifiers* should be assigned to such tiling formats.

Modifier values typically do not *describe* memory layouts. More
precisely, a *modifier*'s lower 56 bits usually have no structure.
Instead, modifiers *name* memory layouts; they name a small set of
vendor-preferred layouts for image sharing. As a consequence, in each
vendor namespace the modifier values are often sequentially allocated
starting at 1.

Each *modifier* is usually supported by a single vendor and its name
matches the pattern `{VENDOR}_FORMAT_MOD_*` or
`DRM_FORMAT_MOD_{VENDOR}_*`. Examples are `I915_FORMAT_MOD_X_TILED` and
`DRM_FORMAT_MOD_BROADCOM_VC4_T_TILED`. An exception is
`DRM_FORMAT_MOD_LINEAR`, which is supported by most vendors.

Many APIs in Linux use *modifiers* to negotiate and specify the memory
layout of shared images. For example, a Wayland compositor and Wayland
client may, by relaying *modifiers* over the Wayland protocol
`zwp_linux_dmabuf_v1`, negotiate a vendor-specific tiling format for a
shared `wl_buffer`. The client may allocate the underlying memory for
the `wl_buffer` with GBM, providing the chosen *modifier* to
`gbm_bo_create_with_modifiers`. The client may then import the
`wl_buffer` into Vulkan for producing image content, providing the
resource’s dma_buf to
[VkImportMemoryFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryFdInfoKHR.html) and its
*modifier* to
[VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html).
The compositor may then import the `wl_buffer` into OpenGL for sampling,
providing the resource’s dma_buf and *modifier* to `eglCreateImage`. The
compositor may also bypass OpenGL and submit the `wl_buffer` directly to
the kernel’s display API, providing the dma_buf and *modifier* through
`drm_mode_fb_cmd2`.

## <a href="#_format_translation" class="anchor"></a>Format Translation

*Modifier*-capable APIs often pair *modifiers* with DRM formats, which
are defined in
[`drm_fourcc.h`](https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/include/uapi/drm/drm_fourcc.h).
However, `VK_EXT_image_drm_format_modifier` uses
[VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) instead of DRM formats. The application must
convert between [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) and DRM format when it sends
or receives a DRM format to or from an external API.

The mapping from [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) to DRM format is lossy.
Therefore, when receiving a DRM format from an external API, often the
application must use information from the external API to accurately map
the DRM format to a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html). For example, DRM formats
do not distinguish between RGB and sRGB (as of 2018-03-28); external
information is required to identify the image’s color space.

The mapping between [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) and DRM format is also
incomplete. For some DRM formats there exist no corresponding Vulkan
format, and for some Vulkan formats there exist no corresponding DRM
format.

## <a href="#_usage_patterns" class="anchor"></a>Usage Patterns

Three primary usage patterns are intended for this extension:

- **Negotiation.** The application negotiates with *modifier*-aware,
  external components to determine sets of image creation parameters
  supported among all components.

  In the Linux ecosystem, the negotiation usually assumes the image is a
  2D, single-sampled, non-mipmapped, non-array image; this extension
  permits that assumption but does not require it. The result of the
  negotiation usually resembles a set of tuples such as *(drmFormat,
  drmFormatModifier)*, where each participating component supports all
  tuples in the set.

  Many details of this negotiation - such as the protocol used during
  negotiation, the set of image creation parameters expressible in the
  protocol, and how the protocol chooses which process and which API
  will create the image - are outside the scope of this specification.

  In this extension,
  [vkGetPhysicalDeviceFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties2.html)
  with
  [VkDrmFormatModifierPropertiesListEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesListEXT.html)
  serves a primary role during the negotiation, and
  [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)
  with
  [VkPhysicalDeviceImageDrmFormatModifierInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageDrmFormatModifierInfoEXT.html)
  serves a secondary role.

- **Import.** The application imports an image with a *modifier*.

  In this pattern, the application receives from an external source the
  image’s memory and its creation parameters, which are often the result
  of the negotiation described above. Some image creation parameters are
  implicitly defined by the external source; for example,
  `VK_IMAGE_TYPE_2D` is often assumed. Some image creation parameters
  are usually explicit, such as the image’s `format`,
  `drmFormatModifier`, and `extent`; and each plane’s `offset` and
  `rowPitch`.

  Before creating the image, the application first verifies that the
  physical device supports the received creation parameters by querying
  [vkGetPhysicalDeviceFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties2.html)
  with
  [VkDrmFormatModifierPropertiesListEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesListEXT.html)
  and
  [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)
  with
  [VkPhysicalDeviceImageDrmFormatModifierInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageDrmFormatModifierInfoEXT.html).
  Then the application creates the image by chaining
  [VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html)
  and
  [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)
  onto [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html).

- **Export.** The application creates an image and allocates its memory.
  Then the application exports to *modifier*-aware consumers the image’s
  memory handles; its creation parameters; its *modifier*; and the
  <a href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout.html" target="_blank"
  rel="noopener"><code>offset</code></a>,
  <a href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout.html" target="_blank"
  rel="noopener"><code>size</code></a>, and
  <a href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout.html" target="_blank"
  rel="noopener"><code>rowPitch</code></a> of each *memory plane*.

  In this pattern, the Vulkan device is the authority for the image; it
  is the allocator of the image’s memory and the decider of the image’s
  creation parameters. When choosing the image’s creation parameters,
  the application usually chooses a tuple *(format, drmFormatModifier)*
  from the result of the negotiation described above. The negotiation’s
  result often contains multiple tuples that share the same format but
  differ in their *modifier*. In this case, the application should defer
  the choice of the image’s *modifier* to the Vulkan implementation by
  providing all such *modifiers* to
  [VkImageDrmFormatModifierListCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierListCreateInfoEXT.html)::`pDrmFormatModifiers`;
  and the implementation should choose from `pDrmFormatModifiers` the
  optimal *modifier* in consideration with the other image parameters.

  The application creates the image by chaining
  [VkImageDrmFormatModifierListCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierListCreateInfoEXT.html)
  and
  [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)
  onto [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html). The protocol and
  APIs by which the application will share the image with external
  consumers will likely determine the value of
  [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)::`handleTypes`.
  The implementation chooses for the image an optimal *modifier* from
  [VkImageDrmFormatModifierListCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierListCreateInfoEXT.html)::`pDrmFormatModifiers`.
  The application then queries the implementation-chosen *modifier* with
  [vkGetImageDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageDrmFormatModifierPropertiesEXT.html),
  and queries the memory layout of each plane with
  [vkGetImageSubresourceLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSubresourceLayout.html).

  The application then allocates the image’s memory with
  [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html), adding chained
  extending structures for external memory; binds it to the image; and
  exports the memory, for example, with
  [vkGetMemoryFdKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryFdKHR.html).

  Finally, the application sends the image’s creation parameters, its
  *modifier*, its per-plane memory layout, and the exported memory
  handle to the external consumers. The details of how the application
  transmits this information to external consumers is outside the scope
  of this specification.

## <a href="#_prior_art" class="anchor"></a>Prior Art

Extension `EGL_EXT_image_dma_buf_import`

[1](https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_image_drm_format_modifier-fn1)^
introduced the ability to create an `EGLImage` by importing for each
plane a dma_buf, offset, and row pitch.

Later, extension `EGL_EXT_image_dma_buf_import_modifiers`

[2](https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_image_drm_format_modifier-fn2)^
introduced the ability to query which combination of formats and
*modifiers* the implementation supports and to specify *modifiers*
during creation of the `EGLImage`.

Extension `EGL_MESA_image_dma_buf_export`

[3](https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_image_drm_format_modifier-fn3)^
is the inverse of `EGL_EXT_image_dma_buf_import_modifiers`.

The Linux kernel modesetting API (KMS), when configuring the display’s
framebuffer with `struct drm_mode_fb_cmd2`

[4](https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_image_drm_format_modifier-fn4)^,
allows one to specify the framebuffer’s *modifier* as well as a
per-plane memory handle, offset, and row pitch.

GBM, a graphics buffer manager for Linux, allows creation of a `gbm_bo`
(that is, a graphics *buffer object*) by importing data similar to that
in `EGL_EXT_image_dma_buf_import_modifiers`

[1](https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_image_drm_format_modifier-fn1)^;
and symmetrically allows exporting the same data from the `gbm_bo`. See
the references to *modifier* and *plane* in `gbm.h`

[5](https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_image_drm_format_modifier-fn5)^.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetImageDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageDrmFormatModifierPropertiesEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesEXT.html)

- [VkImageDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierPropertiesEXT.html)

- Extending [VkFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties2.html):

  - [VkDrmFormatModifierPropertiesListEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesListEXT.html)

- Extending [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html):

  - [VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html)

  - [VkImageDrmFormatModifierListCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierListCreateInfoEXT.html)

- Extending
  [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html):

  - [VkPhysicalDeviceImageDrmFormatModifierInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageDrmFormatModifierInfoEXT.html)

If [VK_KHR_format_feature_flags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_format_feature_flags2.html) or
[Version 1.3](#versions-1.3) is supported:

- [VkDrmFormatModifierProperties2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierProperties2EXT.html)

- Extending [VkFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties2.html):

  - [VkDrmFormatModifierPropertiesList2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesList2EXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_IMAGE_DRM_FORMAT_MODIFIER_EXTENSION_NAME`

- `VK_EXT_IMAGE_DRM_FORMAT_MODIFIER_SPEC_VERSION`

- Extending [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html):

  - `VK_IMAGE_ASPECT_MEMORY_PLANE_0_BIT_EXT`

  - `VK_IMAGE_ASPECT_MEMORY_PLANE_1_BIT_EXT`

  - `VK_IMAGE_ASPECT_MEMORY_PLANE_2_BIT_EXT`

  - `VK_IMAGE_ASPECT_MEMORY_PLANE_3_BIT_EXT`

- Extending [VkImageTiling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageTiling.html):

  - `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`

- Extending [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html):

  - `VK_ERROR_INVALID_DRM_FORMAT_MODIFIER_PLANE_LAYOUT_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DRM_FORMAT_MODIFIER_PROPERTIES_LIST_EXT`

  - `VK_STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_EXPLICIT_CREATE_INFO_EXT`

  - `VK_STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_LIST_CREATE_INFO_EXT`

  - `VK_STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_PROPERTIES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_DRM_FORMAT_MODIFIER_INFO_EXT`

If [VK_KHR_format_feature_flags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_format_feature_flags2.html) or
[Version 1.3](#versions-1.3) is supported:

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DRM_FORMAT_MODIFIER_PROPERTIES_LIST_2_EXT`

## <a href="#_issues" class="anchor"></a>Issues

1\) Should this extension define a single DRM format modifier per
`VkImage`? Or define one per plane?

\+

**RESOLVED**: There exists a single DRM format modifier per `VkImage`.

**DISCUSSION**: Prior art, such as
`EGL_EXT_image_dma_buf_import_modifiers`

[2](https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_image_drm_format_modifier-fn2)^,
`struct drm_mode_fb_cmd2`

[4](https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_image_drm_format_modifier-fn4)^,
and `struct gbm_import_fd_modifier_data`

[5](https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_image_drm_format_modifier-fn5)^,
allows defining one *modifier* per plane. However, developers of the GBM
and kernel APIs concede it was a mistake. Beginning in Linux 4.10, the
kernel requires that the application provide the same DRM format
*modifier* for each plane. (See Linux commit
[bae781b259269590109e8a4a8227331362b88212](https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/commit/?id=bae781b259269590109e8a4a8227331362b88212)).
And GBM provides an entry point, `gbm_bo_get_modifier`, for querying the
*modifier* of the image but does not provide one to query the modifier
of individual planes.

2\) When creating an image with
[VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html),
which is typically used when *importing* an image, should the
application explicitly provide the size of each plane?

\+

**RESOLVED**: No. The application **must** not provide the size. To
enforce this, the API requires that
[VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html)::`pPlaneLayouts->size`
**must** be 0.

**DISCUSSION**: Prior art, such as
`EGL_EXT_image_dma_buf_import_modifiers`

[2](https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_image_drm_format_modifier-fn2)^,
`struct drm_mode_fb_cmd2`

[4](https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_image_drm_format_modifier-fn4)^,
and `struct gbm_import_fd_modifier_data`

[5](https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_image_drm_format_modifier-fn5)^,
omits from the API the size of each plane. Instead, the APIs infer each
plane’s size from the import parameters, which include the image’s pixel
format and a dma_buf, offset, and row pitch for each plane.

However, Vulkan differs from EGL and GBM with regards to image creation
in the following ways:

Differences in Image Creation

- **Undedicated allocation by default.** When importing or exporting a
  set of dma_bufs as an `EGLImage` or `gbm_bo`, common practice mandates
  that each dma_buf’s memory be dedicated (in the sense of
  `VK_KHR_dedicated_allocation`) to the image (though not necessarily
  dedicated to a single plane). In particular, neither the GBM
  documentation nor the EGL extension specifications explicitly state
  this requirement, but in light of common practice this is likely due
  to under-specification rather than intentional omission. In contrast,
  `VK_EXT_image_drm_format_modifier` permits, but does not require, the
  implementation to require dedicated allocations for images created
  with `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`.

- **Separation of image creation and memory allocation.** When importing
  a set of dma_bufs as an `EGLImage` or `gbm_bo`, EGL and GBM create the
  image resource and bind it to memory (the dma_bufs) simultaneously.
  This allows EGL and GBM to query each dma_buf’s size during image
  creation. In Vulkan, image creation and memory allocation are
  independent unless a dedicated allocation is used (as in
  `VK_KHR_dedicated_allocation`). Therefore, without requiring dedicated
  allocation, Vulkan cannot query the size of each dma_buf (or other
  external handle) when calculating the image’s memory layout. Even if
  dedication allocation were required, Vulkan cannot calculate the
  image’s memory layout until after the image is bound to its dma_ufs.

The above differences complicate the potential inference of plane size
in Vulkan. Consider the following problematic cases:

Problematic Plane Size Calculations

- **Padding.** Some plane of the image may require
  implementation-dependent padding.

- **Metadata.** For some *modifiers*, the image may have a metadata
  plane which requires a non-trivial calculation to determine its size.

- **Mipmapped, array, and 3D images.** The implementation may support
  `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT` for images whose
  `mipLevels`, `arrayLayers`, or `depth` is greater than 1. For such
  images with certain *modifiers*, the calculation of each plane’s size
  may be non-trivial.

However, an application-provided plane size solves none of the above
problems.

For simplicity, consider an external image with a single memory plane.
The implementation is obviously capable calculating the image’s size
when its tiling is `VK_IMAGE_TILING_OPTIMAL`. Likewise, any reasonable
implementation is capable of calculating the image’s size when its
tiling uses a supported *modifier*.

Suppose that the external image’s size is smaller than the
implementation-calculated size. If the application provided the external
image’s size to [vkCreateImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImage.html), the implementation
would observe the mismatched size and recognize its inability to
comprehend the external image’s layout (unless the implementation used
the application-provided size to select a refinement of the tiling
layout indicated by the *modifier*, which is strongly discouraged). The
implementation would observe the conflict, and reject image creation
with `VK_ERROR_INVALID_DRM_FORMAT_MODIFIER_PLANE_LAYOUT_EXT`. On the
other hand, if the application did not provide the external image’s size
to [vkCreateImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImage.html), then the application would
observe after calling
[vkGetImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements.html) that
the external image’s size is less than the size required by the
implementation. The application would observe the conflict and refuse to
bind the `VkImage` to the external memory. In both cases, the result is
explicit failure.

Suppose that the external image’s size is larger than the
implementation-calculated size. If the application provided the external
image’s size to [vkCreateImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImage.html), for reasons similar
to above the implementation would observe the mismatched size and
recognize its inability to comprehend the image data residing in the
extra size. The implementation, however, must assume that image data
resides in the entire size provided by the application. The
implementation would observe the conflict and reject image creation with
`VK_ERROR_INVALID_DRM_FORMAT_MODIFIER_PLANE_LAYOUT_EXT`. On the other
hand, if the application did not provide the external image’s size to
[vkCreateImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImage.html), then the application would observe
after calling
[vkGetImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements.html) that
the external image’s size is larger than the implementation-usable size.
The application would observe the conflict and refuse to bind the
`VkImage` to the external memory. In both cases, the result is explicit
failure.

Therefore, an application-provided size provides no benefit, and this
extension should not require it. This decision renders
[VkSubresourceLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout.html)::`size` an unused field
during image creation, and thus introduces a risk that implementations
may require applications to submit sideband creation parameters in the
unused field. To prevent implementations from relying on sideband data,
this extension *requires* the application to set `size` to 0.

### <a href="#_references" class="anchor"></a>References

1.  <span id="VK_EXT_image_drm_format_modifier-fn1"></span>
    [`EGL_EXT_image_dma_buf_import`](https://registry.khronos.org/EGL/extensions/EXT/EGL_EXT_image_dma_buf_import.txt)

2.  <span id="VK_EXT_image_drm_format_modifier-fn2"></span>
    [`EGL_EXT_image_dma_buf_import_modifiers`](https://registry.khronos.org/EGL/extensions/EXT/EGL_EXT_image_dma_buf_import_modifiers.txt)

3.  <span id="VK_EXT_image_drm_format_modifier-fn3"></span>
    [`EGL_MESA_image_dma_buf_export`](https://registry.khronos.org/EGL/extensions/MESA/EGL_MESA_image_dma_buf_export.txt)

4.  <span id="VK_EXT_image_drm_format_modifier-fn4"></span>
    [`struct drm_mode_fb_cmd2`](https://git.kernel.org/cgit/linux/kernel/git/torvalds/linux.git/tree/include/uapi/drm/drm_mode.h?id=refs/tags/v4.10#n392)

5.  <span id="VK_EXT_image_drm_format_modifier-fn5"></span>
    [`gbm.h`](https://cgit.freedesktop.org/mesa/mesa/tree/src/gbm/main/gbm.h?id=refs/tags/mesa-18.0.0-rc1)

### <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-08-29 (Lina Versace)

  - First stable revision

- Revision 2, 2021-09-30 (Jon Leech)

  - Add interaction with
    [`VK_KHR_format_feature_flags2`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_format_feature_flags2.html)
    to `vk.xml`

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_image_drm_format_modifier"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
